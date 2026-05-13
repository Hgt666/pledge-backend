package contract

import (
	"easy-swap/config"
	"easy-swap/logger"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// 全局合约实例
var PledgeInstance *Pledge

func InitContract() {
	// 初始化 PledgeContract
	PledgeContractAddress := common.HexToAddress("0xa0c46D66F17D9F2feC0309881eFe607ea8BF251b")
	client,err := config.DialRPC()
	if err != nil {
		logger.Log.Error("连接以太坊RPC失败: ", zap.Error(err))
	}
	instance,err := NewPledge(PledgeContractAddress,client)
	if err != nil {
		logger.Log.Error("初始化PledgeContract失败: ", zap.Error(err))
	}
	PledgeInstance = instance
}