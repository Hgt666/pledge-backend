// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LibOrderAsset is an auto generated low-level Go binding around an user-defined struct.
type LibOrderAsset struct {
	TokenId    *big.Int
	Collection common.Address
	Amount     *big.Int
}

// LibOrderEditDetail is an auto generated low-level Go binding around an user-defined struct.
type LibOrderEditDetail struct {
	OldOrderKey [32]byte
	NewOrder    LibOrderOrder
}

// LibOrderMatchDetail is an auto generated low-level Go binding around an user-defined struct.
type LibOrderMatchDetail struct {
	SellOrder LibOrderOrder
	BuyOrder  LibOrderOrder
}

// LibOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrder struct {
	Side     uint8
	SaleKind uint8
	Maker    common.Address
	Nft      LibOrderAsset
	Price    *big.Int
	Expiry   uint64
	Salt     uint64
}

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CannotFindNextEmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotFindPrevEmptyKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"}],\"name\":\"CannotInsertDuplicateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInsertEmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInsertExistingKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRemoveEmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRemoveMissingKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msg\",\"type\":\"bytes\"}],\"name\":\"BatchMatchInnerError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"LogCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"name\":\"LogMake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"OrderKey\",\"name\":\"makeOrderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"OrderKey\",\"name\":\"takeOrderKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structLibOrder.Order\",\"name\":\"makeOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structLibOrder.Order\",\"name\":\"takeOrder\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"fillPrice\",\"type\":\"uint128\"}],\"name\":\"LogMatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"OrderKey\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"name\":\"LogSkipOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"newProtocolShare\",\"type\":\"uint128\"}],\"name\":\"LogUpdatedProtocolShare\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msg\",\"type\":\"bytes\"}],\"name\":\"MulticallInnerError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"OrderKey[]\",\"name\":\"orderKeys\",\"type\":\"bytes32[]\"}],\"name\":\"cancelOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"OrderKey\",\"name\":\"oldOrderKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"newOrder\",\"type\":\"tuple\"}],\"internalType\":\"structLibOrder.EditDetail[]\",\"name\":\"editDetails\",\"type\":\"tuple[]\"}],\"name\":\"editOrders\",\"outputs\":[{\"internalType\":\"OrderKey[]\",\"name\":\"newOrderKeys\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"OrderKey\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"}],\"name\":\"getBestOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"orderResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"getBestPrice\",\"outputs\":[{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"}],\"name\":\"getNextBestPrice\",\"outputs\":[{\"internalType\":\"Price\",\"name\":\"nextBestPrice\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"OrderKey\",\"name\":\"firstOrderKey\",\"type\":\"bytes32\"}],\"name\":\"getOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"resultOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"OrderKey\",\"name\":\"nextOrderKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newProtocolShare\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"EIP712Name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"EIP712Version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"newOrders\",\"type\":\"tuple[]\"}],\"name\":\"makeOrders\",\"outputs\":[{\"internalType\":\"OrderKey[]\",\"name\":\"newOrderKeys\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"}],\"name\":\"matchOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"matchOrderWithoutPayback\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"costValue\",\"type\":\"uint128\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"}],\"internalType\":\"structLibOrder.MatchDetail[]\",\"name\":\"matchDetails\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertOnFail\",\"type\":\"bool\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"Price\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"orderQueues\",\"outputs\":[{\"internalType\":\"OrderKey\",\"name\":\"head\",\"type\":\"bytes32\"},{\"internalType\":\"OrderKey\",\"name\":\"tail\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"OrderKey\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structLibOrder.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"Price\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"OrderKey\",\"name\":\"next\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumLibOrder.Side\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"priceTrees\",\"outputs\":[{\"internalType\":\"Price\",\"name\":\"root\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolShare\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newProtocolShare\",\"type\":\"uint128\"}],\"name\":\"setProtocolShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Abi *AbiCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Abi *AbiSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Abi.Contract.Eip712Domain(&_Abi.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Abi *AbiCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Abi.Contract.Eip712Domain(&_Abi.CallOpts)
}

// FilledAmount is a free data retrieval call binding the contract method 0x0b20b7bc.
//
// Solidity: function filledAmount(bytes32 ) view returns(uint256)
func (_Abi *AbiCaller) FilledAmount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "filledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilledAmount is a free data retrieval call binding the contract method 0x0b20b7bc.
//
// Solidity: function filledAmount(bytes32 ) view returns(uint256)
func (_Abi *AbiSession) FilledAmount(arg0 [32]byte) (*big.Int, error) {
	return _Abi.Contract.FilledAmount(&_Abi.CallOpts, arg0)
}

// FilledAmount is a free data retrieval call binding the contract method 0x0b20b7bc.
//
// Solidity: function filledAmount(bytes32 ) view returns(uint256)
func (_Abi *AbiCallerSession) FilledAmount(arg0 [32]byte) (*big.Int, error) {
	return _Abi.Contract.FilledAmount(&_Abi.CallOpts, arg0)
}

// GetBestOrder is a free data retrieval call binding the contract method 0x5dcc7f9c.
//
// Solidity: function getBestOrder(address collection, uint256 tokenId, uint8 side, uint8 saleKind) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) orderResult)
func (_Abi *AbiCaller) GetBestOrder(opts *bind.CallOpts, collection common.Address, tokenId *big.Int, side uint8, saleKind uint8) (LibOrderOrder, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getBestOrder", collection, tokenId, side, saleKind)

	if err != nil {
		return *new(LibOrderOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(LibOrderOrder)).(*LibOrderOrder)

	return out0, err

}

// GetBestOrder is a free data retrieval call binding the contract method 0x5dcc7f9c.
//
// Solidity: function getBestOrder(address collection, uint256 tokenId, uint8 side, uint8 saleKind) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) orderResult)
func (_Abi *AbiSession) GetBestOrder(collection common.Address, tokenId *big.Int, side uint8, saleKind uint8) (LibOrderOrder, error) {
	return _Abi.Contract.GetBestOrder(&_Abi.CallOpts, collection, tokenId, side, saleKind)
}

// GetBestOrder is a free data retrieval call binding the contract method 0x5dcc7f9c.
//
// Solidity: function getBestOrder(address collection, uint256 tokenId, uint8 side, uint8 saleKind) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) orderResult)
func (_Abi *AbiCallerSession) GetBestOrder(collection common.Address, tokenId *big.Int, side uint8, saleKind uint8) (LibOrderOrder, error) {
	return _Abi.Contract.GetBestOrder(&_Abi.CallOpts, collection, tokenId, side, saleKind)
}

// GetBestPrice is a free data retrieval call binding the contract method 0x9e09126c.
//
// Solidity: function getBestPrice(address collection, uint8 side) view returns(uint128 price)
func (_Abi *AbiCaller) GetBestPrice(opts *bind.CallOpts, collection common.Address, side uint8) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getBestPrice", collection, side)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBestPrice is a free data retrieval call binding the contract method 0x9e09126c.
//
// Solidity: function getBestPrice(address collection, uint8 side) view returns(uint128 price)
func (_Abi *AbiSession) GetBestPrice(collection common.Address, side uint8) (*big.Int, error) {
	return _Abi.Contract.GetBestPrice(&_Abi.CallOpts, collection, side)
}

// GetBestPrice is a free data retrieval call binding the contract method 0x9e09126c.
//
// Solidity: function getBestPrice(address collection, uint8 side) view returns(uint128 price)
func (_Abi *AbiCallerSession) GetBestPrice(collection common.Address, side uint8) (*big.Int, error) {
	return _Abi.Contract.GetBestPrice(&_Abi.CallOpts, collection, side)
}

// GetNextBestPrice is a free data retrieval call binding the contract method 0x436a5ed6.
//
// Solidity: function getNextBestPrice(address collection, uint8 side, uint128 price) view returns(uint128 nextBestPrice)
func (_Abi *AbiCaller) GetNextBestPrice(opts *bind.CallOpts, collection common.Address, side uint8, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getNextBestPrice", collection, side, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextBestPrice is a free data retrieval call binding the contract method 0x436a5ed6.
//
// Solidity: function getNextBestPrice(address collection, uint8 side, uint128 price) view returns(uint128 nextBestPrice)
func (_Abi *AbiSession) GetNextBestPrice(collection common.Address, side uint8, price *big.Int) (*big.Int, error) {
	return _Abi.Contract.GetNextBestPrice(&_Abi.CallOpts, collection, side, price)
}

// GetNextBestPrice is a free data retrieval call binding the contract method 0x436a5ed6.
//
// Solidity: function getNextBestPrice(address collection, uint8 side, uint128 price) view returns(uint128 nextBestPrice)
func (_Abi *AbiCallerSession) GetNextBestPrice(collection common.Address, side uint8, price *big.Int) (*big.Int, error) {
	return _Abi.Contract.GetNextBestPrice(&_Abi.CallOpts, collection, side, price)
}

// GetOrders is a free data retrieval call binding the contract method 0x877edbff.
//
// Solidity: function getOrders(address collection, uint256 tokenId, uint8 side, uint8 saleKind, uint256 count, uint128 price, bytes32 firstOrderKey) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] resultOrders, bytes32 nextOrderKey)
func (_Abi *AbiCaller) GetOrders(opts *bind.CallOpts, collection common.Address, tokenId *big.Int, side uint8, saleKind uint8, count *big.Int, price *big.Int, firstOrderKey [32]byte) (struct {
	ResultOrders []LibOrderOrder
	NextOrderKey [32]byte
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getOrders", collection, tokenId, side, saleKind, count, price, firstOrderKey)

	outstruct := new(struct {
		ResultOrders []LibOrderOrder
		NextOrderKey [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResultOrders = *abi.ConvertType(out[0], new([]LibOrderOrder)).(*[]LibOrderOrder)
	outstruct.NextOrderKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetOrders is a free data retrieval call binding the contract method 0x877edbff.
//
// Solidity: function getOrders(address collection, uint256 tokenId, uint8 side, uint8 saleKind, uint256 count, uint128 price, bytes32 firstOrderKey) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] resultOrders, bytes32 nextOrderKey)
func (_Abi *AbiSession) GetOrders(collection common.Address, tokenId *big.Int, side uint8, saleKind uint8, count *big.Int, price *big.Int, firstOrderKey [32]byte) (struct {
	ResultOrders []LibOrderOrder
	NextOrderKey [32]byte
}, error) {
	return _Abi.Contract.GetOrders(&_Abi.CallOpts, collection, tokenId, side, saleKind, count, price, firstOrderKey)
}

// GetOrders is a free data retrieval call binding the contract method 0x877edbff.
//
// Solidity: function getOrders(address collection, uint256 tokenId, uint8 side, uint8 saleKind, uint256 count, uint128 price, bytes32 firstOrderKey) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] resultOrders, bytes32 nextOrderKey)
func (_Abi *AbiCallerSession) GetOrders(collection common.Address, tokenId *big.Int, side uint8, saleKind uint8, count *big.Int, price *big.Int, firstOrderKey [32]byte) (struct {
	ResultOrders []LibOrderOrder
	NextOrderKey [32]byte
}, error) {
	return _Abi.Contract.GetOrders(&_Abi.CallOpts, collection, tokenId, side, saleKind, count, price, firstOrderKey)
}

// OrderQueues is a free data retrieval call binding the contract method 0xe471ba46.
//
// Solidity: function orderQueues(address , uint8 , uint128 ) view returns(bytes32 head, bytes32 tail)
func (_Abi *AbiCaller) OrderQueues(opts *bind.CallOpts, arg0 common.Address, arg1 uint8, arg2 *big.Int) (struct {
	Head [32]byte
	Tail [32]byte
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "orderQueues", arg0, arg1, arg2)

	outstruct := new(struct {
		Head [32]byte
		Tail [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Head = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tail = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// OrderQueues is a free data retrieval call binding the contract method 0xe471ba46.
//
// Solidity: function orderQueues(address , uint8 , uint128 ) view returns(bytes32 head, bytes32 tail)
func (_Abi *AbiSession) OrderQueues(arg0 common.Address, arg1 uint8, arg2 *big.Int) (struct {
	Head [32]byte
	Tail [32]byte
}, error) {
	return _Abi.Contract.OrderQueues(&_Abi.CallOpts, arg0, arg1, arg2)
}

// OrderQueues is a free data retrieval call binding the contract method 0xe471ba46.
//
// Solidity: function orderQueues(address , uint8 , uint128 ) view returns(bytes32 head, bytes32 tail)
func (_Abi *AbiCallerSession) OrderQueues(arg0 common.Address, arg1 uint8, arg2 *big.Int) (struct {
	Head [32]byte
	Tail [32]byte
}, error) {
	return _Abi.Contract.OrderQueues(&_Abi.CallOpts, arg0, arg1, arg2)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) order, bytes32 next)
func (_Abi *AbiCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Order LibOrderOrder
	Next  [32]byte
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Order LibOrderOrder
		Next  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Order = *abi.ConvertType(out[0], new(LibOrderOrder)).(*LibOrderOrder)
	outstruct.Next = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) order, bytes32 next)
func (_Abi *AbiSession) Orders(arg0 [32]byte) (struct {
	Order LibOrderOrder
	Next  [32]byte
}, error) {
	return _Abi.Contract.Orders(&_Abi.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) order, bytes32 next)
func (_Abi *AbiCallerSession) Orders(arg0 [32]byte) (struct {
	Order LibOrderOrder
	Next  [32]byte
}, error) {
	return _Abi.Contract.Orders(&_Abi.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCallerSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiSession) Paused() (bool, error) {
	return _Abi.Contract.Paused(&_Abi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiCallerSession) Paused() (bool, error) {
	return _Abi.Contract.Paused(&_Abi.CallOpts)
}

// PriceTrees is a free data retrieval call binding the contract method 0x7d910dac.
//
// Solidity: function priceTrees(address , uint8 ) view returns(uint128 root)
func (_Abi *AbiCaller) PriceTrees(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "priceTrees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceTrees is a free data retrieval call binding the contract method 0x7d910dac.
//
// Solidity: function priceTrees(address , uint8 ) view returns(uint128 root)
func (_Abi *AbiSession) PriceTrees(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _Abi.Contract.PriceTrees(&_Abi.CallOpts, arg0, arg1)
}

// PriceTrees is a free data retrieval call binding the contract method 0x7d910dac.
//
// Solidity: function priceTrees(address , uint8 ) view returns(uint128 root)
func (_Abi *AbiCallerSession) PriceTrees(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _Abi.Contract.PriceTrees(&_Abi.CallOpts, arg0, arg1)
}

// ProtocolShare is a free data retrieval call binding the contract method 0x1103f315.
//
// Solidity: function protocolShare() view returns(uint128)
func (_Abi *AbiCaller) ProtocolShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "protocolShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolShare is a free data retrieval call binding the contract method 0x1103f315.
//
// Solidity: function protocolShare() view returns(uint128)
func (_Abi *AbiSession) ProtocolShare() (*big.Int, error) {
	return _Abi.Contract.ProtocolShare(&_Abi.CallOpts)
}

// ProtocolShare is a free data retrieval call binding the contract method 0x1103f315.
//
// Solidity: function protocolShare() view returns(uint128)
func (_Abi *AbiCallerSession) ProtocolShare() (*big.Int, error) {
	return _Abi.Contract.ProtocolShare(&_Abi.CallOpts)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderKeys) returns(bool[] successes)
func (_Abi *AbiTransactor) CancelOrders(opts *bind.TransactOpts, orderKeys [][32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "cancelOrders", orderKeys)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderKeys) returns(bool[] successes)
func (_Abi *AbiSession) CancelOrders(orderKeys [][32]byte) (*types.Transaction, error) {
	return _Abi.Contract.CancelOrders(&_Abi.TransactOpts, orderKeys)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x21c77c96.
//
// Solidity: function cancelOrders(bytes32[] orderKeys) returns(bool[] successes)
func (_Abi *AbiTransactorSession) CancelOrders(orderKeys [][32]byte) (*types.Transaction, error) {
	return _Abi.Contract.CancelOrders(&_Abi.TransactOpts, orderKeys)
}

// EditOrders is a paid mutator transaction binding the contract method 0xe8f27e95.
//
// Solidity: function editOrders((bytes32,(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] editDetails) payable returns(bytes32[] newOrderKeys)
func (_Abi *AbiTransactor) EditOrders(opts *bind.TransactOpts, editDetails []LibOrderEditDetail) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "editOrders", editDetails)
}

// EditOrders is a paid mutator transaction binding the contract method 0xe8f27e95.
//
// Solidity: function editOrders((bytes32,(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] editDetails) payable returns(bytes32[] newOrderKeys)
func (_Abi *AbiSession) EditOrders(editDetails []LibOrderEditDetail) (*types.Transaction, error) {
	return _Abi.Contract.EditOrders(&_Abi.TransactOpts, editDetails)
}

// EditOrders is a paid mutator transaction binding the contract method 0xe8f27e95.
//
// Solidity: function editOrders((bytes32,(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] editDetails) payable returns(bytes32[] newOrderKeys)
func (_Abi *AbiTransactorSession) EditOrders(editDetails []LibOrderEditDetail) (*types.Transaction, error) {
	return _Abi.Contract.EditOrders(&_Abi.TransactOpts, editDetails)
}

// Initialize is a paid mutator transaction binding the contract method 0x5d52957e.
//
// Solidity: function initialize(uint128 newProtocolShare, address newVault, string EIP712Name, string EIP712Version) returns()
func (_Abi *AbiTransactor) Initialize(opts *bind.TransactOpts, newProtocolShare *big.Int, newVault common.Address, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "initialize", newProtocolShare, newVault, EIP712Name, EIP712Version)
}

// Initialize is a paid mutator transaction binding the contract method 0x5d52957e.
//
// Solidity: function initialize(uint128 newProtocolShare, address newVault, string EIP712Name, string EIP712Version) returns()
func (_Abi *AbiSession) Initialize(newProtocolShare *big.Int, newVault common.Address, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts, newProtocolShare, newVault, EIP712Name, EIP712Version)
}

// Initialize is a paid mutator transaction binding the contract method 0x5d52957e.
//
// Solidity: function initialize(uint128 newProtocolShare, address newVault, string EIP712Name, string EIP712Version) returns()
func (_Abi *AbiTransactorSession) Initialize(newProtocolShare *big.Int, newVault common.Address, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts, newProtocolShare, newVault, EIP712Name, EIP712Version)
}

// MakeOrders is a paid mutator transaction binding the contract method 0x307725d9.
//
// Solidity: function makeOrders((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] newOrders) payable returns(bytes32[] newOrderKeys)
func (_Abi *AbiTransactor) MakeOrders(opts *bind.TransactOpts, newOrders []LibOrderOrder) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "makeOrders", newOrders)
}

// MakeOrders is a paid mutator transaction binding the contract method 0x307725d9.
//
// Solidity: function makeOrders((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] newOrders) payable returns(bytes32[] newOrderKeys)
func (_Abi *AbiSession) MakeOrders(newOrders []LibOrderOrder) (*types.Transaction, error) {
	return _Abi.Contract.MakeOrders(&_Abi.TransactOpts, newOrders)
}

// MakeOrders is a paid mutator transaction binding the contract method 0x307725d9.
//
// Solidity: function makeOrders((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64)[] newOrders) payable returns(bytes32[] newOrderKeys)
func (_Abi *AbiTransactorSession) MakeOrders(newOrders []LibOrderOrder) (*types.Transaction, error) {
	return _Abi.Contract.MakeOrders(&_Abi.TransactOpts, newOrders)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x882849c9.
//
// Solidity: function matchOrder((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder) payable returns()
func (_Abi *AbiTransactor) MatchOrder(opts *bind.TransactOpts, sellOrder LibOrderOrder, buyOrder LibOrderOrder) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "matchOrder", sellOrder, buyOrder)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x882849c9.
//
// Solidity: function matchOrder((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder) payable returns()
func (_Abi *AbiSession) MatchOrder(sellOrder LibOrderOrder, buyOrder LibOrderOrder) (*types.Transaction, error) {
	return _Abi.Contract.MatchOrder(&_Abi.TransactOpts, sellOrder, buyOrder)
}

// MatchOrder is a paid mutator transaction binding the contract method 0x882849c9.
//
// Solidity: function matchOrder((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder) payable returns()
func (_Abi *AbiTransactorSession) MatchOrder(sellOrder LibOrderOrder, buyOrder LibOrderOrder) (*types.Transaction, error) {
	return _Abi.Contract.MatchOrder(&_Abi.TransactOpts, sellOrder, buyOrder)
}

// MatchOrderWithoutPayback is a paid mutator transaction binding the contract method 0x366c39ed.
//
// Solidity: function matchOrderWithoutPayback((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder, uint256 msgValue) payable returns(uint128 costValue)
func (_Abi *AbiTransactor) MatchOrderWithoutPayback(opts *bind.TransactOpts, sellOrder LibOrderOrder, buyOrder LibOrderOrder, msgValue *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "matchOrderWithoutPayback", sellOrder, buyOrder, msgValue)
}

// MatchOrderWithoutPayback is a paid mutator transaction binding the contract method 0x366c39ed.
//
// Solidity: function matchOrderWithoutPayback((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder, uint256 msgValue) payable returns(uint128 costValue)
func (_Abi *AbiSession) MatchOrderWithoutPayback(sellOrder LibOrderOrder, buyOrder LibOrderOrder, msgValue *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.MatchOrderWithoutPayback(&_Abi.TransactOpts, sellOrder, buyOrder, msgValue)
}

// MatchOrderWithoutPayback is a paid mutator transaction binding the contract method 0x366c39ed.
//
// Solidity: function matchOrderWithoutPayback((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) sellOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) buyOrder, uint256 msgValue) payable returns(uint128 costValue)
func (_Abi *AbiTransactorSession) MatchOrderWithoutPayback(sellOrder LibOrderOrder, buyOrder LibOrderOrder, msgValue *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.MatchOrderWithoutPayback(&_Abi.TransactOpts, sellOrder, buyOrder, msgValue)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xfe971c98.
//
// Solidity: function matchOrders(((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64),(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] matchDetails) payable returns(bool[] successes)
func (_Abi *AbiTransactor) MatchOrders(opts *bind.TransactOpts, matchDetails []LibOrderMatchDetail) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "matchOrders", matchDetails)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xfe971c98.
//
// Solidity: function matchOrders(((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64),(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] matchDetails) payable returns(bool[] successes)
func (_Abi *AbiSession) MatchOrders(matchDetails []LibOrderMatchDetail) (*types.Transaction, error) {
	return _Abi.Contract.MatchOrders(&_Abi.TransactOpts, matchDetails)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xfe971c98.
//
// Solidity: function matchOrders(((uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64),(uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64))[] matchDetails) payable returns(bool[] successes)
func (_Abi *AbiTransactorSession) MatchOrders(matchDetails []LibOrderMatchDetail) (*types.Transaction, error) {
	return _Abi.Contract.MatchOrders(&_Abi.TransactOpts, matchDetails)
}

// Multicall is a paid mutator transaction binding the contract method 0x1e9701d4.
//
// Solidity: function multicall(bytes[] data, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_Abi *AbiTransactor) Multicall(opts *bind.TransactOpts, data [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "multicall", data, revertOnFail)
}

// Multicall is a paid mutator transaction binding the contract method 0x1e9701d4.
//
// Solidity: function multicall(bytes[] data, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_Abi *AbiSession) Multicall(data [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _Abi.Contract.Multicall(&_Abi.TransactOpts, data, revertOnFail)
}

// Multicall is a paid mutator transaction binding the contract method 0x1e9701d4.
//
// Solidity: function multicall(bytes[] data, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_Abi *AbiTransactorSession) Multicall(data [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _Abi.Contract.Multicall(&_Abi.TransactOpts, data, revertOnFail)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactorSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// SetProtocolShare is a paid mutator transaction binding the contract method 0x813e7c2a.
//
// Solidity: function setProtocolShare(uint128 newProtocolShare) returns()
func (_Abi *AbiTransactor) SetProtocolShare(opts *bind.TransactOpts, newProtocolShare *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setProtocolShare", newProtocolShare)
}

// SetProtocolShare is a paid mutator transaction binding the contract method 0x813e7c2a.
//
// Solidity: function setProtocolShare(uint128 newProtocolShare) returns()
func (_Abi *AbiSession) SetProtocolShare(newProtocolShare *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetProtocolShare(&_Abi.TransactOpts, newProtocolShare)
}

// SetProtocolShare is a paid mutator transaction binding the contract method 0x813e7c2a.
//
// Solidity: function setProtocolShare(uint128 newProtocolShare) returns()
func (_Abi *AbiTransactorSession) SetProtocolShare(newProtocolShare *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetProtocolShare(&_Abi.TransactOpts, newProtocolShare)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address newVault) returns()
func (_Abi *AbiTransactor) SetVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setVault", newVault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address newVault) returns()
func (_Abi *AbiSession) SetVault(newVault common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetVault(&_Abi.TransactOpts, newVault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address newVault) returns()
func (_Abi *AbiTransactorSession) SetVault(newVault common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetVault(&_Abi.TransactOpts, newVault)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactorSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_Abi *AbiTransactor) WithdrawETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdrawETH", recipient, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_Abi *AbiSession) WithdrawETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawETH(&_Abi.TransactOpts, recipient, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_Abi *AbiTransactorSession) WithdrawETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.WithdrawETH(&_Abi.TransactOpts, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiSession) Receive() (*types.Transaction, error) {
	return _Abi.Contract.Receive(&_Abi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiTransactorSession) Receive() (*types.Transaction, error) {
	return _Abi.Contract.Receive(&_Abi.TransactOpts)
}

// AbiBatchMatchInnerErrorIterator is returned from FilterBatchMatchInnerError and is used to iterate over the raw logs and unpacked data for BatchMatchInnerError events raised by the Abi contract.
type AbiBatchMatchInnerErrorIterator struct {
	Event *AbiBatchMatchInnerError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiBatchMatchInnerErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiBatchMatchInnerError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiBatchMatchInnerError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiBatchMatchInnerErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiBatchMatchInnerErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiBatchMatchInnerError represents a BatchMatchInnerError event raised by the Abi contract.
type AbiBatchMatchInnerError struct {
	Offset *big.Int
	Msg    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchMatchInnerError is a free log retrieval operation binding the contract event 0x050f709fb65709f10a27682788a9d67fe74de81b310b5c34e3d32f0e2c3ac557.
//
// Solidity: event BatchMatchInnerError(uint256 offset, bytes msg)
func (_Abi *AbiFilterer) FilterBatchMatchInnerError(opts *bind.FilterOpts) (*AbiBatchMatchInnerErrorIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "BatchMatchInnerError")
	if err != nil {
		return nil, err
	}
	return &AbiBatchMatchInnerErrorIterator{contract: _Abi.contract, event: "BatchMatchInnerError", logs: logs, sub: sub}, nil
}

// WatchBatchMatchInnerError is a free log subscription operation binding the contract event 0x050f709fb65709f10a27682788a9d67fe74de81b310b5c34e3d32f0e2c3ac557.
//
// Solidity: event BatchMatchInnerError(uint256 offset, bytes msg)
func (_Abi *AbiFilterer) WatchBatchMatchInnerError(opts *bind.WatchOpts, sink chan<- *AbiBatchMatchInnerError) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "BatchMatchInnerError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiBatchMatchInnerError)
				if err := _Abi.contract.UnpackLog(event, "BatchMatchInnerError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMatchInnerError is a log parse operation binding the contract event 0x050f709fb65709f10a27682788a9d67fe74de81b310b5c34e3d32f0e2c3ac557.
//
// Solidity: event BatchMatchInnerError(uint256 offset, bytes msg)
func (_Abi *AbiFilterer) ParseBatchMatchInnerError(log types.Log) (*AbiBatchMatchInnerError, error) {
	event := new(AbiBatchMatchInnerError)
	if err := _Abi.contract.UnpackLog(event, "BatchMatchInnerError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Abi contract.
type AbiEIP712DomainChangedIterator struct {
	Event *AbiEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiEIP712DomainChanged represents a EIP712DomainChanged event raised by the Abi contract.
type AbiEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Abi *AbiFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AbiEIP712DomainChangedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AbiEIP712DomainChangedIterator{contract: _Abi.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Abi *AbiFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AbiEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiEIP712DomainChanged)
				if err := _Abi.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Abi *AbiFilterer) ParseEIP712DomainChanged(log types.Log) (*AbiEIP712DomainChanged, error) {
	event := new(AbiEIP712DomainChanged)
	if err := _Abi.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Abi contract.
type AbiInitializedIterator struct {
	Event *AbiInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInitialized represents a Initialized event raised by the Abi contract.
type AbiInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) FilterInitialized(opts *bind.FilterOpts) (*AbiInitializedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AbiInitializedIterator{contract: _Abi.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AbiInitialized) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInitialized)
				if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) ParseInitialized(log types.Log) (*AbiInitialized, error) {
	event := new(AbiInitialized)
	if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLogCancelIterator is returned from FilterLogCancel and is used to iterate over the raw logs and unpacked data for LogCancel events raised by the Abi contract.
type AbiLogCancelIterator struct {
	Event *AbiLogCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLogCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLogCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLogCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLogCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogCancel represents a LogCancel event raised by the Abi contract.
type AbiLogCancel struct {
	OrderKey [32]byte
	Maker    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogCancel is a free log retrieval operation binding the contract event 0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd.
//
// Solidity: event LogCancel(bytes32 indexed orderKey, address indexed maker)
func (_Abi *AbiFilterer) FilterLogCancel(opts *bind.FilterOpts, orderKey [][32]byte, maker []common.Address) (*AbiLogCancelIterator, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "LogCancel", orderKeyRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &AbiLogCancelIterator{contract: _Abi.contract, event: "LogCancel", logs: logs, sub: sub}, nil
}

// WatchLogCancel is a free log subscription operation binding the contract event 0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd.
//
// Solidity: event LogCancel(bytes32 indexed orderKey, address indexed maker)
func (_Abi *AbiFilterer) WatchLogCancel(opts *bind.WatchOpts, sink chan<- *AbiLogCancel, orderKey [][32]byte, maker []common.Address) (event.Subscription, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "LogCancel", orderKeyRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogCancel)
				if err := _Abi.contract.UnpackLog(event, "LogCancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogCancel is a log parse operation binding the contract event 0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd.
//
// Solidity: event LogCancel(bytes32 indexed orderKey, address indexed maker)
func (_Abi *AbiFilterer) ParseLogCancel(log types.Log) (*AbiLogCancel, error) {
	event := new(AbiLogCancel)
	if err := _Abi.contract.UnpackLog(event, "LogCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLogMakeIterator is returned from FilterLogMake and is used to iterate over the raw logs and unpacked data for LogMake events raised by the Abi contract.
type AbiLogMakeIterator struct {
	Event *AbiLogMake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLogMakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogMake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLogMake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLogMakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLogMakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogMake represents a LogMake event raised by the Abi contract.
type AbiLogMake struct {
	OrderKey [32]byte
	Side     uint8
	SaleKind uint8
	Maker    common.Address
	Nft      LibOrderAsset
	Price    *big.Int
	Expiry   uint64
	Salt     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogMake is a free log retrieval operation binding the contract event 0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6.
//
// Solidity: event LogMake(bytes32 orderKey, uint8 indexed side, uint8 indexed saleKind, address indexed maker, (uint256,address,uint96) nft, uint128 price, uint64 expiry, uint64 salt)
func (_Abi *AbiFilterer) FilterLogMake(opts *bind.FilterOpts, side []uint8, saleKind []uint8, maker []common.Address) (*AbiLogMakeIterator, error) {

	var sideRule []interface{}
	for _, sideItem := range side {
		sideRule = append(sideRule, sideItem)
	}
	var saleKindRule []interface{}
	for _, saleKindItem := range saleKind {
		saleKindRule = append(saleKindRule, saleKindItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "LogMake", sideRule, saleKindRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &AbiLogMakeIterator{contract: _Abi.contract, event: "LogMake", logs: logs, sub: sub}, nil
}

// WatchLogMake is a free log subscription operation binding the contract event 0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6.
//
// Solidity: event LogMake(bytes32 orderKey, uint8 indexed side, uint8 indexed saleKind, address indexed maker, (uint256,address,uint96) nft, uint128 price, uint64 expiry, uint64 salt)
func (_Abi *AbiFilterer) WatchLogMake(opts *bind.WatchOpts, sink chan<- *AbiLogMake, side []uint8, saleKind []uint8, maker []common.Address) (event.Subscription, error) {

	var sideRule []interface{}
	for _, sideItem := range side {
		sideRule = append(sideRule, sideItem)
	}
	var saleKindRule []interface{}
	for _, saleKindItem := range saleKind {
		saleKindRule = append(saleKindRule, saleKindItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "LogMake", sideRule, saleKindRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogMake)
				if err := _Abi.contract.UnpackLog(event, "LogMake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMake is a log parse operation binding the contract event 0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6.
//
// Solidity: event LogMake(bytes32 orderKey, uint8 indexed side, uint8 indexed saleKind, address indexed maker, (uint256,address,uint96) nft, uint128 price, uint64 expiry, uint64 salt)
func (_Abi *AbiFilterer) ParseLogMake(log types.Log) (*AbiLogMake, error) {
	event := new(AbiLogMake)
	if err := _Abi.contract.UnpackLog(event, "LogMake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLogMatchIterator is returned from FilterLogMatch and is used to iterate over the raw logs and unpacked data for LogMatch events raised by the Abi contract.
type AbiLogMatchIterator struct {
	Event *AbiLogMatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLogMatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogMatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLogMatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLogMatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLogMatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogMatch represents a LogMatch event raised by the Abi contract.
type AbiLogMatch struct {
	MakeOrderKey [32]byte
	TakeOrderKey [32]byte
	MakeOrder    LibOrderOrder
	TakeOrder    LibOrderOrder
	FillPrice    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogMatch is a free log retrieval operation binding the contract event 0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e.
//
// Solidity: event LogMatch(bytes32 indexed makeOrderKey, bytes32 indexed takeOrderKey, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) makeOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) takeOrder, uint128 fillPrice)
func (_Abi *AbiFilterer) FilterLogMatch(opts *bind.FilterOpts, makeOrderKey [][32]byte, takeOrderKey [][32]byte) (*AbiLogMatchIterator, error) {

	var makeOrderKeyRule []interface{}
	for _, makeOrderKeyItem := range makeOrderKey {
		makeOrderKeyRule = append(makeOrderKeyRule, makeOrderKeyItem)
	}
	var takeOrderKeyRule []interface{}
	for _, takeOrderKeyItem := range takeOrderKey {
		takeOrderKeyRule = append(takeOrderKeyRule, takeOrderKeyItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "LogMatch", makeOrderKeyRule, takeOrderKeyRule)
	if err != nil {
		return nil, err
	}
	return &AbiLogMatchIterator{contract: _Abi.contract, event: "LogMatch", logs: logs, sub: sub}, nil
}

// WatchLogMatch is a free log subscription operation binding the contract event 0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e.
//
// Solidity: event LogMatch(bytes32 indexed makeOrderKey, bytes32 indexed takeOrderKey, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) makeOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) takeOrder, uint128 fillPrice)
func (_Abi *AbiFilterer) WatchLogMatch(opts *bind.WatchOpts, sink chan<- *AbiLogMatch, makeOrderKey [][32]byte, takeOrderKey [][32]byte) (event.Subscription, error) {

	var makeOrderKeyRule []interface{}
	for _, makeOrderKeyItem := range makeOrderKey {
		makeOrderKeyRule = append(makeOrderKeyRule, makeOrderKeyItem)
	}
	var takeOrderKeyRule []interface{}
	for _, takeOrderKeyItem := range takeOrderKey {
		takeOrderKeyRule = append(takeOrderKeyRule, takeOrderKeyItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "LogMatch", makeOrderKeyRule, takeOrderKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogMatch)
				if err := _Abi.contract.UnpackLog(event, "LogMatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogMatch is a log parse operation binding the contract event 0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e.
//
// Solidity: event LogMatch(bytes32 indexed makeOrderKey, bytes32 indexed takeOrderKey, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) makeOrder, (uint8,uint8,address,(uint256,address,uint96),uint128,uint64,uint64) takeOrder, uint128 fillPrice)
func (_Abi *AbiFilterer) ParseLogMatch(log types.Log) (*AbiLogMatch, error) {
	event := new(AbiLogMatch)
	if err := _Abi.contract.UnpackLog(event, "LogMatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLogSkipOrderIterator is returned from FilterLogSkipOrder and is used to iterate over the raw logs and unpacked data for LogSkipOrder events raised by the Abi contract.
type AbiLogSkipOrderIterator struct {
	Event *AbiLogSkipOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLogSkipOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogSkipOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLogSkipOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLogSkipOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLogSkipOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogSkipOrder represents a LogSkipOrder event raised by the Abi contract.
type AbiLogSkipOrder struct {
	OrderKey [32]byte
	Salt     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSkipOrder is a free log retrieval operation binding the contract event 0x43d1f368251ebe03c021962d50212d072e7ccee5c8ad3f541d93d0dc43bbd420.
//
// Solidity: event LogSkipOrder(bytes32 orderKey, uint64 salt)
func (_Abi *AbiFilterer) FilterLogSkipOrder(opts *bind.FilterOpts) (*AbiLogSkipOrderIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "LogSkipOrder")
	if err != nil {
		return nil, err
	}
	return &AbiLogSkipOrderIterator{contract: _Abi.contract, event: "LogSkipOrder", logs: logs, sub: sub}, nil
}

// WatchLogSkipOrder is a free log subscription operation binding the contract event 0x43d1f368251ebe03c021962d50212d072e7ccee5c8ad3f541d93d0dc43bbd420.
//
// Solidity: event LogSkipOrder(bytes32 orderKey, uint64 salt)
func (_Abi *AbiFilterer) WatchLogSkipOrder(opts *bind.WatchOpts, sink chan<- *AbiLogSkipOrder) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "LogSkipOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogSkipOrder)
				if err := _Abi.contract.UnpackLog(event, "LogSkipOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogSkipOrder is a log parse operation binding the contract event 0x43d1f368251ebe03c021962d50212d072e7ccee5c8ad3f541d93d0dc43bbd420.
//
// Solidity: event LogSkipOrder(bytes32 orderKey, uint64 salt)
func (_Abi *AbiFilterer) ParseLogSkipOrder(log types.Log) (*AbiLogSkipOrder, error) {
	event := new(AbiLogSkipOrder)
	if err := _Abi.contract.UnpackLog(event, "LogSkipOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLogUpdatedProtocolShareIterator is returned from FilterLogUpdatedProtocolShare and is used to iterate over the raw logs and unpacked data for LogUpdatedProtocolShare events raised by the Abi contract.
type AbiLogUpdatedProtocolShareIterator struct {
	Event *AbiLogUpdatedProtocolShare // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLogUpdatedProtocolShareIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogUpdatedProtocolShare)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLogUpdatedProtocolShare)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLogUpdatedProtocolShareIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLogUpdatedProtocolShareIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogUpdatedProtocolShare represents a LogUpdatedProtocolShare event raised by the Abi contract.
type AbiLogUpdatedProtocolShare struct {
	NewProtocolShare *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogUpdatedProtocolShare is a free log retrieval operation binding the contract event 0x0b52884d4590055c8791518c34458834f75feed662340ef0b5af2898e7a9be9f.
//
// Solidity: event LogUpdatedProtocolShare(uint128 indexed newProtocolShare)
func (_Abi *AbiFilterer) FilterLogUpdatedProtocolShare(opts *bind.FilterOpts, newProtocolShare []*big.Int) (*AbiLogUpdatedProtocolShareIterator, error) {

	var newProtocolShareRule []interface{}
	for _, newProtocolShareItem := range newProtocolShare {
		newProtocolShareRule = append(newProtocolShareRule, newProtocolShareItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "LogUpdatedProtocolShare", newProtocolShareRule)
	if err != nil {
		return nil, err
	}
	return &AbiLogUpdatedProtocolShareIterator{contract: _Abi.contract, event: "LogUpdatedProtocolShare", logs: logs, sub: sub}, nil
}

// WatchLogUpdatedProtocolShare is a free log subscription operation binding the contract event 0x0b52884d4590055c8791518c34458834f75feed662340ef0b5af2898e7a9be9f.
//
// Solidity: event LogUpdatedProtocolShare(uint128 indexed newProtocolShare)
func (_Abi *AbiFilterer) WatchLogUpdatedProtocolShare(opts *bind.WatchOpts, sink chan<- *AbiLogUpdatedProtocolShare, newProtocolShare []*big.Int) (event.Subscription, error) {

	var newProtocolShareRule []interface{}
	for _, newProtocolShareItem := range newProtocolShare {
		newProtocolShareRule = append(newProtocolShareRule, newProtocolShareItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "LogUpdatedProtocolShare", newProtocolShareRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogUpdatedProtocolShare)
				if err := _Abi.contract.UnpackLog(event, "LogUpdatedProtocolShare", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUpdatedProtocolShare is a log parse operation binding the contract event 0x0b52884d4590055c8791518c34458834f75feed662340ef0b5af2898e7a9be9f.
//
// Solidity: event LogUpdatedProtocolShare(uint128 indexed newProtocolShare)
func (_Abi *AbiFilterer) ParseLogUpdatedProtocolShare(log types.Log) (*AbiLogUpdatedProtocolShare, error) {
	event := new(AbiLogUpdatedProtocolShare)
	if err := _Abi.contract.UnpackLog(event, "LogUpdatedProtocolShare", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLogWithdrawETHIterator is returned from FilterLogWithdrawETH and is used to iterate over the raw logs and unpacked data for LogWithdrawETH events raised by the Abi contract.
type AbiLogWithdrawETHIterator struct {
	Event *AbiLogWithdrawETH // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLogWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLogWithdrawETH)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLogWithdrawETH)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLogWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLogWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLogWithdrawETH represents a LogWithdrawETH event raised by the Abi contract.
type AbiLogWithdrawETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawETH is a free log retrieval operation binding the contract event 0xab0f46966ebb6f17d26b8f6add4624e39fcfdeaccfeaba589d57f81cb5cb6666.
//
// Solidity: event LogWithdrawETH(address recipient, uint256 amount)
func (_Abi *AbiFilterer) FilterLogWithdrawETH(opts *bind.FilterOpts) (*AbiLogWithdrawETHIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "LogWithdrawETH")
	if err != nil {
		return nil, err
	}
	return &AbiLogWithdrawETHIterator{contract: _Abi.contract, event: "LogWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawETH is a free log subscription operation binding the contract event 0xab0f46966ebb6f17d26b8f6add4624e39fcfdeaccfeaba589d57f81cb5cb6666.
//
// Solidity: event LogWithdrawETH(address recipient, uint256 amount)
func (_Abi *AbiFilterer) WatchLogWithdrawETH(opts *bind.WatchOpts, sink chan<- *AbiLogWithdrawETH) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "LogWithdrawETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLogWithdrawETH)
				if err := _Abi.contract.UnpackLog(event, "LogWithdrawETH", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogWithdrawETH is a log parse operation binding the contract event 0xab0f46966ebb6f17d26b8f6add4624e39fcfdeaccfeaba589d57f81cb5cb6666.
//
// Solidity: event LogWithdrawETH(address recipient, uint256 amount)
func (_Abi *AbiFilterer) ParseLogWithdrawETH(log types.Log) (*AbiLogWithdrawETH, error) {
	event := new(AbiLogWithdrawETH)
	if err := _Abi.contract.UnpackLog(event, "LogWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMulticallInnerErrorIterator is returned from FilterMulticallInnerError and is used to iterate over the raw logs and unpacked data for MulticallInnerError events raised by the Abi contract.
type AbiMulticallInnerErrorIterator struct {
	Event *AbiMulticallInnerError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiMulticallInnerErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMulticallInnerError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiMulticallInnerError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiMulticallInnerErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMulticallInnerErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMulticallInnerError represents a MulticallInnerError event raised by the Abi contract.
type AbiMulticallInnerError struct {
	Offset *big.Int
	Msg    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMulticallInnerError is a free log retrieval operation binding the contract event 0x91837fc535a3750e88fdd9cb575731d49f1cc5c44d57bd4501659912cb2d5941.
//
// Solidity: event MulticallInnerError(uint256 offset, bytes msg)
func (_Abi *AbiFilterer) FilterMulticallInnerError(opts *bind.FilterOpts) (*AbiMulticallInnerErrorIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "MulticallInnerError")
	if err != nil {
		return nil, err
	}
	return &AbiMulticallInnerErrorIterator{contract: _Abi.contract, event: "MulticallInnerError", logs: logs, sub: sub}, nil
}

// WatchMulticallInnerError is a free log subscription operation binding the contract event 0x91837fc535a3750e88fdd9cb575731d49f1cc5c44d57bd4501659912cb2d5941.
//
// Solidity: event MulticallInnerError(uint256 offset, bytes msg)
func (_Abi *AbiFilterer) WatchMulticallInnerError(opts *bind.WatchOpts, sink chan<- *AbiMulticallInnerError) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "MulticallInnerError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMulticallInnerError)
				if err := _Abi.contract.UnpackLog(event, "MulticallInnerError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMulticallInnerError is a log parse operation binding the contract event 0x91837fc535a3750e88fdd9cb575731d49f1cc5c44d57bd4501659912cb2d5941.
//
// Solidity: event MulticallInnerError(uint256 offset, bytes msg)
func (_Abi *AbiFilterer) ParseMulticallInnerError(log types.Log) (*AbiMulticallInnerError, error) {
	event := new(AbiMulticallInnerError)
	if err := _Abi.contract.UnpackLog(event, "MulticallInnerError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abi contract.
type AbiOwnershipTransferredIterator struct {
	Event *AbiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOwnershipTransferred represents a OwnershipTransferred event raised by the Abi contract.
type AbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbiOwnershipTransferredIterator{contract: _Abi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOwnershipTransferred)
				if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) ParseOwnershipTransferred(log types.Log) (*AbiOwnershipTransferred, error) {
	event := new(AbiOwnershipTransferred)
	if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Abi contract.
type AbiPausedIterator struct {
	Event *AbiPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPaused represents a Paused event raised by the Abi contract.
type AbiPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) FilterPaused(opts *bind.FilterOpts) (*AbiPausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AbiPausedIterator{contract: _Abi.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AbiPaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPaused)
				if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) ParsePaused(log types.Log) (*AbiPaused, error) {
	event := new(AbiPaused)
	if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Abi contract.
type AbiUnpausedIterator struct {
	Event *AbiUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUnpaused represents a Unpaused event raised by the Abi contract.
type AbiUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AbiUnpausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AbiUnpausedIterator{contract: _Abi.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AbiUnpaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUnpaused)
				if err := _Abi.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) ParseUnpaused(log types.Log) (*AbiUnpaused, error) {
	event := new(AbiUnpaused)
	if err := _Abi.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
