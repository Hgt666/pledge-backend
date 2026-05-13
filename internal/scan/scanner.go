package scan

import (
	"context"
	"easy-swap/config"
	"easy-swap/internal/parser"
	"easy-swap/logger"
	"easy-swap/model"

	// "easy-swap/model"
	"easy-swap/service"
	// "encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// 全局扫链实例
var Global *Scanner

// 注入业务Service
var (
	scanSvc = service.NewScanService()
	pledgeSvc = service.NewPledgeService()
)

// 事件签名常量
const (
	DepositLendSig   = "0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea"
	DepositBorrowSig = "0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5"
	StateChange        = "0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb"
	ClaimLendSig     = "0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd"
	ClaimBorrow      = "0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd"
	// Finish           = "0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb"
	WithdrawLend     = "0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d"
)

// Scanner 扫链服务结构体
type Scanner struct {
	client    *ethclient.Client
	lastBlock uint64
	blockHash string
	contracts []common.Address
}

// NewScanner 创建扫链实例
func NewScanner() (*Scanner, error) {
	client, err := config.DialRPC()
	if err != nil {
		return nil, err
	}

	// 初始化扫描高度
	scanSvc.Init()
	lastBlock := scanSvc.GetLastBlock()

	s := &Scanner{
		client:    client,
		lastBlock: lastBlock,
		contracts: []common.Address{config.PledgeContractAddr},
	}
	Global = s
	return s, nil
}

// Start 启动定时扫链
func (s *Scanner) Start(ctx context.Context) error {
	logger.Log.Info("✅ 扫链服务启动", zap.Uint64("lastblock", s.lastBlock))
	// ticker := time.NewTicker(time.Duration(config.GlobalConfig.Chain.ScanInterval) * time.Second)
	interval := time.Second * 3
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 每3秒执行一次
			_ = s.handleForkRollback()
			if err := s.scanBlockLoop(); err != nil {
				logger.Log.Error("❌ 扫描异常,重连RPC", zap.Error(err))
				_ = s.reconnectRPC()
			}

		case <-ctx.Done():
			// 优雅退出（企业必须有）
			// log.Println("🛑 扫链服务停止")
			logger.Log.Warn("🛑 扫链服务停止")

		}
	}
}

// scanBlockLoop 扫描区间区块日志
func (s *Scanner) scanBlockLoop() error {
	ctx := context.Background()
	latestBlock, err := s.client.BlockNumber(ctx)
	if err != nil {
		return err
	}
	// 🔥 修复：免费RPC限制，每次最多扫 5 个区块
	toBlock := s.lastBlock + 5
	// fmt.Printf("toBlock:%d\n",toBlock)
	fmt.Printf("链上最新的区块：%d\n", latestBlock)
	fmt.Printf("本地最新的区块：%d\n", s.lastBlock)
	if toBlock > latestBlock {
		toBlock = latestBlock
	}

	// 批量查询区块日志
	logs, err := s.client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: s.contracts,
		FromBlock: big.NewInt(int64(s.lastBlock + 1)),
		ToBlock:   big.NewInt(int64(toBlock)),
	})
	// fmt.Println(logs)
	if err != nil {
		return err
	}

	// 遍历解析事件
	for _, vLog := range logs {
		s.dispatchEvent(&vLog)
	}

	// 更新扫描高度
	s.lastBlock = toBlock
	// 通过lastBlock获取区块哈希
	block, err := s.client.BlockByNumber(ctx, big.NewInt(int64(s.lastBlock)))
	if err != nil {
		logger.Log.Error("获取区块信息失败", zap.Uint64("block", s.lastBlock), zap.Error(err))
		return err
	}
	blockHash := block.Hash().Hex()
	s.blockHash = blockHash
	scanSvc.SaveLastBlock(s.lastBlock, s.blockHash)
	return nil
}

// handleForkRollback 检测链分叉并回滚脏数据
func (s *Scanner) handleForkRollback() error {
	ctx := context.Background()
	latestOnChain, err := s.client.BlockNumber(ctx)
	if err != nil {
		return err
	}

	// ====== 【1】先判断：本地高度 > 链上高度 → 分叉 ======
	if s.lastBlock > latestOnChain {
		rollbackPoint := latestOnChain - config.GlobalConfig.Chain.SafeBlock
		logger.Log.Warn("⚠️ 检测到链分叉（本地高度 > 链上高度）", zap.Uint64("rollback", rollbackPoint))

		// 回滚业务数据

		// 重置高度
		s.lastBlock = rollbackPoint
		scanSvc.SaveLastBlock(s.lastBlock, s.blockHash)
		return nil
	}

	// ====== 【2】关键：判断 同高度但区块哈希不匹配 → 分叉（最常见！）======
	// 只有本地已经扫描过区块，才做哈希校验
	if s.lastBlock <= 0 {
		return nil // 第一次启动，不校验
	}

	// 获取【链上】当前区块的哈希
	onChainBlock, err := s.client.BlockByNumber(ctx, big.NewInt(int64(s.lastBlock)))
	if err != nil {
		return err
	}
	chainBlockHash := onChainBlock.Hash().Hex()

	// 获取【本地数据库】存储的当前区块哈希
	localBlockHash, err := scanSvc.GetLastBlockHash() // 你需要实现这个方法
	if err != nil {
		// 🟢 本地哈希为空 → 第一次扫描，不分叉，直接返回
		logger.Log.Info("🟢 首次扫描，不进行分叉校验", zap.Uint64("block", s.lastBlock))
		return nil
	}

	// 🔴 区块哈希不一致 → 真正的链分叉！（高度一样，块不一样）
	if localBlockHash != chainBlockHash {
		rollbackPoint := s.lastBlock - config.GlobalConfig.Chain.SafeBlock
		logger.Log.Warn("⚠️ 检测到链分叉（区块哈希不匹配）",
			zap.Uint64("block", s.lastBlock),
			zap.String("local_hash", localBlockHash),
			zap.String("chain_hash", chainBlockHash),
			zap.Uint64("rollback_to", rollbackPoint),
		)

		// 回滚脏数据

		// 重置扫描高度
		s.lastBlock = rollbackPoint
		scanSvc.SaveLastBlock(s.lastBlock, s.blockHash)
	}

	return nil
}

// ReScan 手动区块回溯重扫
func (s *Scanner) ReScan(step uint64) {
	if s.lastBlock <= step {
		s.lastBlock = 0
	} else {
		s.lastBlock -= step
	}
	scanSvc.SaveLastBlock(s.lastBlock, s.blockHash)
	logger.Log.Info("🔄 手动回溯成功，当前扫描起始区块：%d", zap.Uint64("lastBlock", s.lastBlock))
}

// dispatchEvent 事件分发器
func (s *Scanner) dispatchEvent(log *types.Log) {
	sig := log.Topics[0].Hex()
	// fmt.Printf("sig:%s\n",sig)
	switch sig {
	case DepositLendSig:
		// logger.Log.Info("📥 监测到 DepositLend 事件", zap.Uint64("block", log.BlockNumber), zap.String("txHash", log.TxHash.Hex()))
		event,err :=parser.ParseDepositLend(log)
		if err != nil {
			return
		}
		userDepositLend := &model.UserDeposit{
			ChainID : config.GlobalConfig.Chain.ChainID,
			PoolID : "0",
			UserAddress : event.From.Hex(),
			TxHash : log.TxHash.Hex(),
			BlockNumber : log.BlockNumber,
			Amount : event.Amount.String(),
			Interest : "0",
			Duration : 0,
			IsClaimed : false,
			Status : 0,
		}
		err = pledgeSvc.SaveDepositLend(userDepositLend)
		if err != nil {
			logger.Log.Error("处理 DepositLend 事件失败", zap.Error(err))
		}

		// 更新借贷池总存款（可选，视业务需求而定）
		


	case DepositBorrowSig:
		// logger.Log.Info("📥 监测到 DepositBorrow 事件", zap.Uint64("block", log.BlockNumber), zap.String("txHash", log.TxHash.Hex()))
		event, err := parser.ParseDepositBorrow(log)
		if err != nil {
			return
		}
		userDepositBorrow := &model.UserPledgeLoan{
			ChainID : config.GlobalConfig.Chain.ChainID,
			PoolID : "0",
			UserAddress : event.From.Hex(),
			TxHash : log.TxHash.Hex(),
			BlockNumber : log.BlockNumber,
			Duration : 0,
			CollateralAmount : event.Amount.String(),
			LoanAmount : event.MintAmount.String(),
			InterestAmount : "0",
			HealthRate : 0,
			Status : 0,
			Version : 1,
		}
		err = pledgeSvc.SaveDepositBorrow(userDepositBorrow)
		if err != nil {
			logger.Log.Error("处理 DepositBorrow 事件失败", zap.Error(err))
		}

	case StateChange:
		// logger.Log.Info("📥 监测到 StateChange 事件", zap.Uint64("block", log.BlockNumber), zap.String("txHash", log.TxHash.Hex()))
		event, err := parser.ParseStateChange(log)
		if err != nil {
			return
		}
		lendingPool := &model.LendingPool{
			ChainID: config.GlobalConfig.Chain.ChainID,
			PoolID: event.Pid.String(),
			MultiSignID: 1,
			PoolAddress: "0",
			Name: "0",
			CollateralToken: "0",
			LoanToken: "0",
			Duration: uint64(event.AfterState.Uint64()),
			LTV: 0,
			LiquidationRate: 0,
			ARP: 0,
			TotalDeposit: "0",
			TotalLoan: "0",
			Status: uint(event.AfterState.Uint64()),
		}
		err = pledgeSvc.SaveLendingPool(lendingPool)
		if err != nil {
			logger.Log.Error("处理 StateChange 事件失败", zap.Error(err))
		}

	case ClaimLendSig:
		// logger.Log.Info("📥 监测到 ClaimLend 事件", zap.Uint64("block", log.BlockNumber), zap.String("txHash", log.TxHash.Hex()))

	case ClaimBorrow:
		// logger.Log.Info("📥 监测到 ClaimBorrow 事件", zap.Uint64("block", log.BlockNumber), zap.String("txHash", log.TxHash.Hex()))

	case WithdrawLend:
		// logger.Log.Info("📥 监测到 WithdrawLend 事件", zap.Uint64("block", log.BlockNumber), zap.String("txHash", log.TxHash.Hex()))
	default:
		// logger.Log.Warn("未识别的事件", zap.String("sig", sig), zap.Uint64("block", log.BlockNumber))

	}
}

// reconnectRPC 重连节点
func (s *Scanner) reconnectRPC() error {
	client, err := config.DialRPC()
	if err != nil {
		return err
	}
	s.client = client
	logger.Log.Info("✅ RPC节点重连成功")
	return nil
}
