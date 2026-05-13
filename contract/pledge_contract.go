package contract

import (
	"context"
	"easy-swap/utils"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 合约地址（你自己的）
const PledgePoolAddress = "0xa0c46D66F17D9F2feC0309881eFe607ea8BF251b"

// PledgeContract 合约结构体
type PledgeContract struct {
	client     *ethclient.Client
	abi        *abi.ABI
	contractAd common.Address
}

// NewPledgeContract 初始化合约
func NewPledgeContract(client *ethclient.Client) (*PledgeContract, error) {
	// 读取 ABI
	abiBytes, err := utils.ReadABIFile("PledgePool.abi")
	if err != nil {
		return nil, err
	}

	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}

	return &PledgeContract{
		client:     client,
		abi:        &parsedABI,
		contractAd: common.HexToAddress(PledgePoolAddress),
	}, nil
}

// ===================== 核心：解析 Deposit 事件 =====================
type DepositEvent struct {
	PoolId    string
	User      common.Address
	Amount    *big.Int
	TxHash    common.Hash
	BlockNum  uint64
}

// ParseDepositEvents 解析某个块的 Deposit 事件
func (p *PledgeContract) ParseDepositEvents(blockNumber uint64) ([]DepositEvent, error) {
	// 1. 过滤查询
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		ToBlock:   big.NewInt(int64(blockNumber)),
		Addresses: []common.Address{p.contractAd},
		Topics:    [][]common.Hash{{p.abi.Events["Deposit"].ID}},
	}

	// 2. 获取链上日志
	logs, err := p.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var events []DepositEvent
	for _, log := range logs {
		var e DepositEvent
		// 3. 解析日志到结构体
		err := p.abi.UnpackIntoInterface(&e, "Deposit", log.Data)
		if err != nil {
			continue
		}
		e.User = common.BytesToAddress(log.Topics[1].Bytes())
		e.TxHash = log.TxHash
		e.BlockNum = log.BlockNumber
		events = append(events, e)
	}

	return events, nil
}