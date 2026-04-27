package service

import (
	"easy-swap/config"
	"easy-swap/internal/contract"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NFTChainService struct{}

func NewNFTChainService() *NFTChainService {
	return &NFTChainService{}
}

func (s *NFTChainService) GetNFTInfo() (name string, symbol string, err error) {
	client, err := config.DialRPC()
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	name, err = contract.GetName(client)
	if err != nil {
		return "", "", err
	}

	symbol, _ = contract.GetSymbol(client)
	return name, symbol, nil
}

func (s *NFTChainService) GetOwner(tokenId uint64) (string, error) {
	client, err := config.DialRPC()
	if err != nil {
		return "", err
	}
	defer client.Close()

	owner, err := contract.OwnerOf(client, big.NewInt(int64(tokenId)))
	if err != nil {
		return "", err
	}
	return owner.Hex(), nil
}

func (s *NFTChainService) GetBalance(addr string) (string, error) {
	client, err := config.DialRPC()
	if err != nil {
		return "", err
	}
	defer client.Close()

	balance, err := contract.BalanceOf(client, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}