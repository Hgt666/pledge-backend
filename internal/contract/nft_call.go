package contract

import (
	"context"
	"easy-swap/config"
	"easy-swap/utils"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var nftABI = `[{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

var parsedABI, _ = abi.JSON(strings.NewReader(nftABI))

// GetName 获取NFT名称
func GetName(client *ethclient.Client) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), utils.RPCTimeout)
	defer cancel()

	data, err := parsedABI.Pack("name")
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		To:   &config.NftContractAddr,
		Data: data,
	}

	res, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", err
	}

	var name string
	err = parsedABI.UnpackIntoInterface(&name, "name", res)
	return name, err
}

// GetSymbol 获取NFT符号
func GetSymbol(client *ethclient.Client) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), utils.RPCTimeout)
	defer cancel()
	data, _ := parsedABI.Pack("symbol")
	msg := ethereum.CallMsg{To: &config.NftContractAddr, Data: data}
	res, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", err
	}
	var symbol string
	_ = parsedABI.UnpackIntoInterface(&symbol, "symbol", res)
	return symbol, err
}

// OwnerOf 获取拥有者
func OwnerOf(client *ethclient.Client, tokenId *big.Int) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), utils.RPCTimeout)
	defer cancel()
	data, _ := parsedABI.Pack("ownerOf", tokenId)
	msg := ethereum.CallMsg{To: &config.NftContractAddr, Data: data}
	res, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return common.Address{}, err
	}
	var owner common.Address
	_ = parsedABI.UnpackIntoInterface(&owner, "ownerOf", res)
	return owner, err
}

// BalanceOf 获取余额
func BalanceOf(client *ethclient.Client, owner common.Address) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), utils.RPCTimeout)
	defer cancel()
	data, _ := parsedABI.Pack("balanceOf", owner)
	msg := ethereum.CallMsg{To: &config.NftContractAddr, Data: data}
	res, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	var balance *big.Int
	_ = parsedABI.UnpackIntoInterface(&balance, "balanceOf", res)
	return balance, err
}