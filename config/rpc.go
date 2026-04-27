package config

import (
	"context"
	"easy-swap/utils"

	"github.com/ethereum/go-ethereum/ethclient"
)

// DialRPC 连接以太坊RPC节点
func DialRPC() (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), utils.RPCTimeout)
	defer cancel()
	return ethclient.DialContext(ctx, GlobalConfig.Chain.RPC)
}