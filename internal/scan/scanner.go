package scan

import (
	"context"
	"easy-swap/config"
	"easy-swap/internal/parser"
	"easy-swap/logger"
	"easy-swap/model"
	"easy-swap/service"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
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
	scanSvc  = service.NewScanService()
	nftSvc   = service.NewNFTService()
	orderSvc = service.NewOrderService()
)

// 事件签名常量
const (
	TransferSig       = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	OrderCreatedSig   = "0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6"
	OrderCancelledSig = "0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd"
	OrderFilledSig    = "0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e"
)

// Scanner 扫链服务结构体
type Scanner struct {
	client    *ethclient.Client
	lastBlock uint64
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
		contracts: []common.Address{config.NftContractAddr, config.MarketContractAddr},
	}
	Global = s
	return s, nil
}

// Start 启动定时扫链
func (s *Scanner) Start() {
	logger.Log.Info("✅ 扫链服务启动",zap.Uint64("block",s.lastBlock))
	// ticker := time.NewTicker(time.Duration(config.GlobalConfig.Chain.ScanInterval) * time.Second)
	interval := time.Second * 3
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 监听系统退出信号（企业必备）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	

	// ✅ 这就是你要背的标准写法
	for {
		select {
		case <-ticker.C:
			// 每3秒执行一次
			_ = s.handleForkRollback()
			if err := s.scanBlockLoop(); err != nil {
				logger.Log.Error("❌ 扫描异常，重连RPC",zap.Error(err))
				_ = s.reconnectRPC()
			}

		case <-quit:
			// 优雅退出（企业必须有）
			// log.Println("🛑 扫链服务停止")
			logger.Log.Warn("🛑 扫链服务停止")
			return
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
	if toBlock > latestBlock {
		toBlock = latestBlock
	}

	// 批量查询区块日志
	logs, err := s.client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: s.contracts,
		FromBlock: big.NewInt(int64(s.lastBlock + 1)),
		ToBlock:   big.NewInt(int64(toBlock)),
	})
	if err != nil {
		return err
	}

	// 遍历解析事件
	for _, vLog := range logs {
		s.dispatchEvent(&vLog)
	}

	// 更新扫描高度
	s.lastBlock = latestBlock
	scanSvc.SaveLastBlock(s.lastBlock)
	return nil
}

// handleForkRollback 检测链分叉并回滚脏数据
func (s *Scanner) handleForkRollback() error {
	ctx := context.Background()
	latestOnChain, err := s.client.BlockNumber(ctx)
	if err != nil {
		return err
	}

	// 本地高度大于链上高度，说明出现区块分叉
	if s.lastBlock > latestOnChain {
		rollbackPoint := latestOnChain - config.GlobalConfig.Chain.SafeBlock
		// log.Printf("⚠️ 检测到链分叉，数据回滚至区块：%d", rollbackPoint)
		logger.Log.Warn("⚠️ 检测到链分叉，数据回滚至区块:",zap.Uint64("rollback",rollbackPoint))

		// 业务层回滚数据
		_ = nftSvc.RollbackByBlock(rollbackPoint)
		_ = orderSvc.RollbackByBlock(rollbackPoint)

		// 重置扫描高度
		s.lastBlock = rollbackPoint
		scanSvc.SaveLastBlock(s.lastBlock)
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
	scanSvc.SaveLastBlock(s.lastBlock)
	logger.Log.Info("🔄 手动回溯成功，当前扫描起始区块：%d",zap.Uint64("lastBlock",s.lastBlock))
}

// dispatchEvent 事件分发器
func (s *Scanner) dispatchEvent(log *types.Log) {
	sig := log.Topics[0].Hex()
	fmt.Printf("sig:%s\n",sig)
	switch sig {
	case TransferSig:
		ev, err := parser.ParseTransfer(log)
		if err != nil {
			return
		}
		_ = nftSvc.SaveTransfer(&model.NftTransfer{
			TokenID:     ev.TokenId,
			FromAddress: ev.From.Hex(),
			ToAddress:   ev.To.Hex(),
			TxHash:      log.TxHash.Hex(),
			BlockNumber: log.BlockNumber,
		})

	case OrderCreatedSig:
		ev, err := parser.ParseOrderCreated(log)
		if err != nil {
			return
		}
		_ = orderSvc.CreateOrder(&model.MarketOrder{
			OrderId:     ev.OrderId,
			TokenID:     ev.TokenId,
			Seller:      ev.Seller.Hex(),
			Price:       ev.Price,
			Status:      0,
			TxHash:      log.TxHash.Hex(),
			BlockNumber: log.BlockNumber,
		})

	case OrderCancelledSig:
		ev, err := parser.ParseOrderCancelled(log)
		if err != nil {
			return
		}
		_ = orderSvc.CancelOrder(ev.OrderId)

	case OrderFilledSig:
		ev, err := parser.ParseOrderFilled(log)
		if err != nil {
			return
		}
		_ = orderSvc.FillOrder(ev.OrderId)
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