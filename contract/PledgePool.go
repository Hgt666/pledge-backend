// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// PledgeMetaData contains all meta data concerning the Pledge contract.
var PledgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_multiSignature\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimLend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"DepositBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"DepositLend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyBorrowWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyLendWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recieptor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"RefundBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"RefundLend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLendFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newBorrowFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeAddress\",\"type\":\"address\"}],\"name\":\"SetFeeAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinAmount\",\"type\":\"uint256\"}],\"name\":\"SetMinAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSwapAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSwapAddress\",\"type\":\"address\"}],\"name\":\"SetSwapRouterAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"beforeState\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"afterState\",\"type\":\"uint256\"}],\"name\":\"StateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromCoin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toCoin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toValue\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawLend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"borrowFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"checkoutFinish\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"checkoutLiquidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"checkoutSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"claimBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"claimLend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_settleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_interestRate\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_martgageRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lendToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrowToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_jpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_autoLiquidateThreshold\",\"type\":\"uint256\"}],\"name\":\"createPoolInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"depositBorrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"depositLend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyBorrowWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyLendWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"finish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMultiSignatureAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getPoolState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingPriceView\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIBscPledgeOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolBaseInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"settleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lendSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"martgageRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lendToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowToken\",\"type\":\"address\"},{\"internalType\":\"enumPledgePool.PoolState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"contractIDebtToken\",\"name\":\"spCoin\",\"type\":\"address\"},{\"internalType\":\"contractIDebtToken\",\"name\":\"jpCoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"autoLiquidateThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolDataInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"settleAmountLend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settleAmountBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishAmountLend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishAmountBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmounLend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmounBorrow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"refundBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"refundLend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lendFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userBorrowInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasNoRefund\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasNoClaim\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userLendInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasNoRefund\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasNoClaim\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_jpAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawLend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PledgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PledgeMetaData.ABI instead.
var PledgeABI = PledgeMetaData.ABI

// Pledge is an auto generated Go binding around an Ethereum contract.
type Pledge struct {
	PledgeCaller     // Read-only binding to the contract
	PledgeTransactor // Write-only binding to the contract
	PledgeFilterer   // Log filterer for contract events
}

// PledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgeSession struct {
	Contract     *Pledge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgeCallerSession struct {
	Contract *PledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgeTransactorSession struct {
	Contract     *PledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgeRaw struct {
	Contract *Pledge // Generic contract binding to access the raw methods on
}

// PledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgeCallerRaw struct {
	Contract *PledgeCaller // Generic read-only contract binding to access the raw methods on
}

// PledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgeTransactorRaw struct {
	Contract *PledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledge creates a new instance of Pledge, bound to a specific deployed contract.
func NewPledge(address common.Address, backend bind.ContractBackend) (*Pledge, error) {
	contract, err := bindPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pledge{PledgeCaller: PledgeCaller{contract: contract}, PledgeTransactor: PledgeTransactor{contract: contract}, PledgeFilterer: PledgeFilterer{contract: contract}}, nil
}

// NewPledgeCaller creates a new read-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeCaller(address common.Address, caller bind.ContractCaller) (*PledgeCaller, error) {
	contract, err := bindPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeCaller{contract: contract}, nil
}

// NewPledgeTransactor creates a new write-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgeTransactor, error) {
	contract, err := bindPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeTransactor{contract: contract}, nil
}

// NewPledgeFilterer creates a new log filterer instance of Pledge, bound to a specific deployed contract.
func NewPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgeFilterer, error) {
	contract, err := bindPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgeFilterer{contract: contract}, nil
}

// bindPledge binds a generic wrapper to an already deployed contract.
func bindPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PledgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.PledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transact(opts, method, params...)
}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_Pledge *PledgeCaller) BorrowFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "borrowFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_Pledge *PledgeSession) BorrowFee() (*big.Int, error) {
	return _Pledge.Contract.BorrowFee(&_Pledge.CallOpts)
}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_Pledge *PledgeCallerSession) BorrowFee() (*big.Int, error) {
	return _Pledge.Contract.BorrowFee(&_Pledge.CallOpts)
}

// CheckoutFinish is a free data retrieval call binding the contract method 0x6abd7f29.
//
// Solidity: function checkoutFinish(uint256 _pid) view returns(bool)
func (_Pledge *PledgeCaller) CheckoutFinish(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "checkoutFinish", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckoutFinish is a free data retrieval call binding the contract method 0x6abd7f29.
//
// Solidity: function checkoutFinish(uint256 _pid) view returns(bool)
func (_Pledge *PledgeSession) CheckoutFinish(_pid *big.Int) (bool, error) {
	return _Pledge.Contract.CheckoutFinish(&_Pledge.CallOpts, _pid)
}

// CheckoutFinish is a free data retrieval call binding the contract method 0x6abd7f29.
//
// Solidity: function checkoutFinish(uint256 _pid) view returns(bool)
func (_Pledge *PledgeCallerSession) CheckoutFinish(_pid *big.Int) (bool, error) {
	return _Pledge.Contract.CheckoutFinish(&_Pledge.CallOpts, _pid)
}

// CheckoutLiquidate is a free data retrieval call binding the contract method 0x08e7305f.
//
// Solidity: function checkoutLiquidate(uint256 _pid) view returns(bool)
func (_Pledge *PledgeCaller) CheckoutLiquidate(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "checkoutLiquidate", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckoutLiquidate is a free data retrieval call binding the contract method 0x08e7305f.
//
// Solidity: function checkoutLiquidate(uint256 _pid) view returns(bool)
func (_Pledge *PledgeSession) CheckoutLiquidate(_pid *big.Int) (bool, error) {
	return _Pledge.Contract.CheckoutLiquidate(&_Pledge.CallOpts, _pid)
}

// CheckoutLiquidate is a free data retrieval call binding the contract method 0x08e7305f.
//
// Solidity: function checkoutLiquidate(uint256 _pid) view returns(bool)
func (_Pledge *PledgeCallerSession) CheckoutLiquidate(_pid *big.Int) (bool, error) {
	return _Pledge.Contract.CheckoutLiquidate(&_Pledge.CallOpts, _pid)
}

// CheckoutSettle is a free data retrieval call binding the contract method 0x14c090cc.
//
// Solidity: function checkoutSettle(uint256 _pid) view returns(bool)
func (_Pledge *PledgeCaller) CheckoutSettle(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "checkoutSettle", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckoutSettle is a free data retrieval call binding the contract method 0x14c090cc.
//
// Solidity: function checkoutSettle(uint256 _pid) view returns(bool)
func (_Pledge *PledgeSession) CheckoutSettle(_pid *big.Int) (bool, error) {
	return _Pledge.Contract.CheckoutSettle(&_Pledge.CallOpts, _pid)
}

// CheckoutSettle is a free data retrieval call binding the contract method 0x14c090cc.
//
// Solidity: function checkoutSettle(uint256 _pid) view returns(bool)
func (_Pledge *PledgeCallerSession) CheckoutSettle(_pid *big.Int) (bool, error) {
	return _Pledge.Contract.CheckoutSettle(&_Pledge.CallOpts, _pid)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Pledge *PledgeCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Pledge *PledgeSession) FeeAddress() (common.Address, error) {
	return _Pledge.Contract.FeeAddress(&_Pledge.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Pledge *PledgeCallerSession) FeeAddress() (common.Address, error) {
	return _Pledge.Contract.FeeAddress(&_Pledge.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_Pledge *PledgeCaller) GetMultiSignatureAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getMultiSignatureAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_Pledge *PledgeSession) GetMultiSignatureAddress() (common.Address, error) {
	return _Pledge.Contract.GetMultiSignatureAddress(&_Pledge.CallOpts)
}

// GetMultiSignatureAddress is a free data retrieval call binding the contract method 0x638c7e17.
//
// Solidity: function getMultiSignatureAddress() view returns(address)
func (_Pledge *PledgeCallerSession) GetMultiSignatureAddress() (common.Address, error) {
	return _Pledge.Contract.GetMultiSignatureAddress(&_Pledge.CallOpts)
}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_Pledge *PledgeCaller) GetPoolState(opts *bind.CallOpts, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getPoolState", _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_Pledge *PledgeSession) GetPoolState(_pid *big.Int) (*big.Int, error) {
	return _Pledge.Contract.GetPoolState(&_Pledge.CallOpts, _pid)
}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_Pledge *PledgeCallerSession) GetPoolState(_pid *big.Int) (*big.Int, error) {
	return _Pledge.Contract.GetPoolState(&_Pledge.CallOpts, _pid)
}

// GetUnderlyingPriceView is a free data retrieval call binding the contract method 0xc9333756.
//
// Solidity: function getUnderlyingPriceView(uint256 _pid) view returns(uint256[2])
func (_Pledge *PledgeCaller) GetUnderlyingPriceView(opts *bind.CallOpts, _pid *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUnderlyingPriceView", _pid)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetUnderlyingPriceView is a free data retrieval call binding the contract method 0xc9333756.
//
// Solidity: function getUnderlyingPriceView(uint256 _pid) view returns(uint256[2])
func (_Pledge *PledgeSession) GetUnderlyingPriceView(_pid *big.Int) ([2]*big.Int, error) {
	return _Pledge.Contract.GetUnderlyingPriceView(&_Pledge.CallOpts, _pid)
}

// GetUnderlyingPriceView is a free data retrieval call binding the contract method 0xc9333756.
//
// Solidity: function getUnderlyingPriceView(uint256 _pid) view returns(uint256[2])
func (_Pledge *PledgeCallerSession) GetUnderlyingPriceView(_pid *big.Int) ([2]*big.Int, error) {
	return _Pledge.Contract.GetUnderlyingPriceView(&_Pledge.CallOpts, _pid)
}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_Pledge *PledgeCaller) GlobalPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "globalPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_Pledge *PledgeSession) GlobalPaused() (bool, error) {
	return _Pledge.Contract.GlobalPaused(&_Pledge.CallOpts)
}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_Pledge *PledgeCallerSession) GlobalPaused() (bool, error) {
	return _Pledge.Contract.GlobalPaused(&_Pledge.CallOpts)
}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_Pledge *PledgeCaller) LendFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "lendFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_Pledge *PledgeSession) LendFee() (*big.Int, error) {
	return _Pledge.Contract.LendFee(&_Pledge.CallOpts)
}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_Pledge *PledgeCallerSession) LendFee() (*big.Int, error) {
	return _Pledge.Contract.LendFee(&_Pledge.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Pledge *PledgeCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Pledge *PledgeSession) MinAmount() (*big.Int, error) {
	return _Pledge.Contract.MinAmount(&_Pledge.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Pledge *PledgeCallerSession) MinAmount() (*big.Int, error) {
	return _Pledge.Contract.MinAmount(&_Pledge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Pledge *PledgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Pledge *PledgeSession) Oracle() (common.Address, error) {
	return _Pledge.Contract.Oracle(&_Pledge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Pledge *PledgeCallerSession) Oracle() (common.Address, error) {
	return _Pledge.Contract.Oracle(&_Pledge.CallOpts)
}

// PoolBaseInfo is a free data retrieval call binding the contract method 0x5a5a971e.
//
// Solidity: function poolBaseInfo(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 martgageRate, address lendToken, address borrowToken, uint8 state, address spCoin, address jpCoin, uint256 autoLiquidateThreshold)
func (_Pledge *PledgeCaller) PoolBaseInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MartgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	SpCoin                 common.Address
	JpCoin                 common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "poolBaseInfo", arg0)

	outstruct := new(struct {
		SettleTime             *big.Int
		EndTime                *big.Int
		InterestRate           *big.Int
		MaxSupply              *big.Int
		LendSupply             *big.Int
		BorrowSupply           *big.Int
		MartgageRate           *big.Int
		LendToken              common.Address
		BorrowToken            common.Address
		State                  uint8
		SpCoin                 common.Address
		JpCoin                 common.Address
		AutoLiquidateThreshold *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SettleTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InterestRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LendSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BorrowSupply = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MartgageRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LendToken = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.BorrowToken = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.SpCoin = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.JpCoin = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)
	outstruct.AutoLiquidateThreshold = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolBaseInfo is a free data retrieval call binding the contract method 0x5a5a971e.
//
// Solidity: function poolBaseInfo(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 martgageRate, address lendToken, address borrowToken, uint8 state, address spCoin, address jpCoin, uint256 autoLiquidateThreshold)
func (_Pledge *PledgeSession) PoolBaseInfo(arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MartgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	SpCoin                 common.Address
	JpCoin                 common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	return _Pledge.Contract.PoolBaseInfo(&_Pledge.CallOpts, arg0)
}

// PoolBaseInfo is a free data retrieval call binding the contract method 0x5a5a971e.
//
// Solidity: function poolBaseInfo(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 martgageRate, address lendToken, address borrowToken, uint8 state, address spCoin, address jpCoin, uint256 autoLiquidateThreshold)
func (_Pledge *PledgeCallerSession) PoolBaseInfo(arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MartgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	SpCoin                 common.Address
	JpCoin                 common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	return _Pledge.Contract.PoolBaseInfo(&_Pledge.CallOpts, arg0)
}

// PoolDataInfo is a free data retrieval call binding the contract method 0x0177b68c.
//
// Solidity: function poolDataInfo(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmounLend, uint256 liquidationAmounBorrow)
func (_Pledge *PledgeCaller) PoolDataInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SettleAmountLend       *big.Int
	SettleAmountBorrow     *big.Int
	FinishAmountLend       *big.Int
	FinishAmountBorrow     *big.Int
	LiquidationAmounLend   *big.Int
	LiquidationAmounBorrow *big.Int
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "poolDataInfo", arg0)

	outstruct := new(struct {
		SettleAmountLend       *big.Int
		SettleAmountBorrow     *big.Int
		FinishAmountLend       *big.Int
		FinishAmountBorrow     *big.Int
		LiquidationAmounLend   *big.Int
		LiquidationAmounBorrow *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SettleAmountLend = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SettleAmountBorrow = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinishAmountLend = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FinishAmountBorrow = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LiquidationAmounLend = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidationAmounBorrow = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolDataInfo is a free data retrieval call binding the contract method 0x0177b68c.
//
// Solidity: function poolDataInfo(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmounLend, uint256 liquidationAmounBorrow)
func (_Pledge *PledgeSession) PoolDataInfo(arg0 *big.Int) (struct {
	SettleAmountLend       *big.Int
	SettleAmountBorrow     *big.Int
	FinishAmountLend       *big.Int
	FinishAmountBorrow     *big.Int
	LiquidationAmounLend   *big.Int
	LiquidationAmounBorrow *big.Int
}, error) {
	return _Pledge.Contract.PoolDataInfo(&_Pledge.CallOpts, arg0)
}

// PoolDataInfo is a free data retrieval call binding the contract method 0x0177b68c.
//
// Solidity: function poolDataInfo(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmounLend, uint256 liquidationAmounBorrow)
func (_Pledge *PledgeCallerSession) PoolDataInfo(arg0 *big.Int) (struct {
	SettleAmountLend       *big.Int
	SettleAmountBorrow     *big.Int
	FinishAmountLend       *big.Int
	FinishAmountBorrow     *big.Int
	LiquidationAmounLend   *big.Int
	LiquidationAmounBorrow *big.Int
}, error) {
	return _Pledge.Contract.PoolDataInfo(&_Pledge.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Pledge *PledgeCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Pledge *PledgeSession) PoolLength() (*big.Int, error) {
	return _Pledge.Contract.PoolLength(&_Pledge.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Pledge *PledgeCallerSession) PoolLength() (*big.Int, error) {
	return _Pledge.Contract.PoolLength(&_Pledge.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Pledge *PledgeCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Pledge *PledgeSession) SwapRouter() (common.Address, error) {
	return _Pledge.Contract.SwapRouter(&_Pledge.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Pledge *PledgeCallerSession) SwapRouter() (common.Address, error) {
	return _Pledge.Contract.SwapRouter(&_Pledge.CallOpts)
}

// UserBorrowInfo is a free data retrieval call binding the contract method 0x3c9fadc3.
//
// Solidity: function userBorrowInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_Pledge *PledgeCaller) UserBorrowInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "userBorrowInfo", arg0, arg1)

	outstruct := new(struct {
		StakeAmount  *big.Int
		RefundAmount *big.Int
		HasNoRefund  bool
		HasNoClaim   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefundAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasNoRefund = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.HasNoClaim = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserBorrowInfo is a free data retrieval call binding the contract method 0x3c9fadc3.
//
// Solidity: function userBorrowInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_Pledge *PledgeSession) UserBorrowInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _Pledge.Contract.UserBorrowInfo(&_Pledge.CallOpts, arg0, arg1)
}

// UserBorrowInfo is a free data retrieval call binding the contract method 0x3c9fadc3.
//
// Solidity: function userBorrowInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_Pledge *PledgeCallerSession) UserBorrowInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _Pledge.Contract.UserBorrowInfo(&_Pledge.CallOpts, arg0, arg1)
}

// UserLendInfo is a free data retrieval call binding the contract method 0xbb176a64.
//
// Solidity: function userLendInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_Pledge *PledgeCaller) UserLendInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "userLendInfo", arg0, arg1)

	outstruct := new(struct {
		StakeAmount  *big.Int
		RefundAmount *big.Int
		HasNoRefund  bool
		HasNoClaim   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefundAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasNoRefund = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.HasNoClaim = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserLendInfo is a free data retrieval call binding the contract method 0xbb176a64.
//
// Solidity: function userLendInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_Pledge *PledgeSession) UserLendInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _Pledge.Contract.UserLendInfo(&_Pledge.CallOpts, arg0, arg1)
}

// UserLendInfo is a free data retrieval call binding the contract method 0xbb176a64.
//
// Solidity: function userLendInfo(address , uint256 ) view returns(uint256 stakeAmount, uint256 refundAmount, bool hasNoRefund, bool hasNoClaim)
func (_Pledge *PledgeCallerSession) UserLendInfo(arg0 common.Address, arg1 *big.Int) (struct {
	StakeAmount  *big.Int
	RefundAmount *big.Int
	HasNoRefund  bool
	HasNoClaim   bool
}, error) {
	return _Pledge.Contract.UserLendInfo(&_Pledge.CallOpts, arg0, arg1)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) ClaimBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "claimBorrow", _pid)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_Pledge *PledgeSession) ClaimBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.ClaimBorrow(&_Pledge.TransactOpts, _pid)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) ClaimBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.ClaimBorrow(&_Pledge.TransactOpts, _pid)
}

// ClaimLend is a paid mutator transaction binding the contract method 0x6c42fed2.
//
// Solidity: function claimLend(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) ClaimLend(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "claimLend", _pid)
}

// ClaimLend is a paid mutator transaction binding the contract method 0x6c42fed2.
//
// Solidity: function claimLend(uint256 _pid) returns()
func (_Pledge *PledgeSession) ClaimLend(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.ClaimLend(&_Pledge.TransactOpts, _pid)
}

// ClaimLend is a paid mutator transaction binding the contract method 0x6c42fed2.
//
// Solidity: function claimLend(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) ClaimLend(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.ClaimLend(&_Pledge.TransactOpts, _pid)
}

// CreatePoolInfo is a paid mutator transaction binding the contract method 0xebded017.
//
// Solidity: function createPoolInfo(uint256 _settleTime, uint256 _endTime, uint64 _interestRate, uint256 _maxSupply, uint256 _martgageRate, address _lendToken, address _borrowToken, address _spToken, address _jpToken, uint256 _autoLiquidateThreshold) returns()
func (_Pledge *PledgeTransactor) CreatePoolInfo(opts *bind.TransactOpts, _settleTime *big.Int, _endTime *big.Int, _interestRate uint64, _maxSupply *big.Int, _martgageRate *big.Int, _lendToken common.Address, _borrowToken common.Address, _spToken common.Address, _jpToken common.Address, _autoLiquidateThreshold *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "createPoolInfo", _settleTime, _endTime, _interestRate, _maxSupply, _martgageRate, _lendToken, _borrowToken, _spToken, _jpToken, _autoLiquidateThreshold)
}

// CreatePoolInfo is a paid mutator transaction binding the contract method 0xebded017.
//
// Solidity: function createPoolInfo(uint256 _settleTime, uint256 _endTime, uint64 _interestRate, uint256 _maxSupply, uint256 _martgageRate, address _lendToken, address _borrowToken, address _spToken, address _jpToken, uint256 _autoLiquidateThreshold) returns()
func (_Pledge *PledgeSession) CreatePoolInfo(_settleTime *big.Int, _endTime *big.Int, _interestRate uint64, _maxSupply *big.Int, _martgageRate *big.Int, _lendToken common.Address, _borrowToken common.Address, _spToken common.Address, _jpToken common.Address, _autoLiquidateThreshold *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.CreatePoolInfo(&_Pledge.TransactOpts, _settleTime, _endTime, _interestRate, _maxSupply, _martgageRate, _lendToken, _borrowToken, _spToken, _jpToken, _autoLiquidateThreshold)
}

// CreatePoolInfo is a paid mutator transaction binding the contract method 0xebded017.
//
// Solidity: function createPoolInfo(uint256 _settleTime, uint256 _endTime, uint64 _interestRate, uint256 _maxSupply, uint256 _martgageRate, address _lendToken, address _borrowToken, address _spToken, address _jpToken, uint256 _autoLiquidateThreshold) returns()
func (_Pledge *PledgeTransactorSession) CreatePoolInfo(_settleTime *big.Int, _endTime *big.Int, _interestRate uint64, _maxSupply *big.Int, _martgageRate *big.Int, _lendToken common.Address, _borrowToken common.Address, _spToken common.Address, _jpToken common.Address, _autoLiquidateThreshold *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.CreatePoolInfo(&_Pledge.TransactOpts, _settleTime, _endTime, _interestRate, _maxSupply, _martgageRate, _lendToken, _borrowToken, _spToken, _jpToken, _autoLiquidateThreshold)
}

// DepositBorrow is a paid mutator transaction binding the contract method 0x16f941b5.
//
// Solidity: function depositBorrow(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_Pledge *PledgeTransactor) DepositBorrow(opts *bind.TransactOpts, _pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "depositBorrow", _pid, _stakeAmount)
}

// DepositBorrow is a paid mutator transaction binding the contract method 0x16f941b5.
//
// Solidity: function depositBorrow(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_Pledge *PledgeSession) DepositBorrow(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.DepositBorrow(&_Pledge.TransactOpts, _pid, _stakeAmount)
}

// DepositBorrow is a paid mutator transaction binding the contract method 0x16f941b5.
//
// Solidity: function depositBorrow(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_Pledge *PledgeTransactorSession) DepositBorrow(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.DepositBorrow(&_Pledge.TransactOpts, _pid, _stakeAmount)
}

// DepositLend is a paid mutator transaction binding the contract method 0x90590da0.
//
// Solidity: function depositLend(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_Pledge *PledgeTransactor) DepositLend(opts *bind.TransactOpts, _pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "depositLend", _pid, _stakeAmount)
}

// DepositLend is a paid mutator transaction binding the contract method 0x90590da0.
//
// Solidity: function depositLend(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_Pledge *PledgeSession) DepositLend(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.DepositLend(&_Pledge.TransactOpts, _pid, _stakeAmount)
}

// DepositLend is a paid mutator transaction binding the contract method 0x90590da0.
//
// Solidity: function depositLend(uint256 _pid, uint256 _stakeAmount) payable returns()
func (_Pledge *PledgeTransactorSession) DepositLend(_pid *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.DepositLend(&_Pledge.TransactOpts, _pid, _stakeAmount)
}

// EmergencyBorrowWithdrawal is a paid mutator transaction binding the contract method 0xe271fa0c.
//
// Solidity: function emergencyBorrowWithdrawal(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) EmergencyBorrowWithdrawal(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "emergencyBorrowWithdrawal", _pid)
}

// EmergencyBorrowWithdrawal is a paid mutator transaction binding the contract method 0xe271fa0c.
//
// Solidity: function emergencyBorrowWithdrawal(uint256 _pid) returns()
func (_Pledge *PledgeSession) EmergencyBorrowWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.EmergencyBorrowWithdrawal(&_Pledge.TransactOpts, _pid)
}

// EmergencyBorrowWithdrawal is a paid mutator transaction binding the contract method 0xe271fa0c.
//
// Solidity: function emergencyBorrowWithdrawal(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) EmergencyBorrowWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.EmergencyBorrowWithdrawal(&_Pledge.TransactOpts, _pid)
}

// EmergencyLendWithdrawal is a paid mutator transaction binding the contract method 0xbf38b8f6.
//
// Solidity: function emergencyLendWithdrawal(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) EmergencyLendWithdrawal(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "emergencyLendWithdrawal", _pid)
}

// EmergencyLendWithdrawal is a paid mutator transaction binding the contract method 0xbf38b8f6.
//
// Solidity: function emergencyLendWithdrawal(uint256 _pid) returns()
func (_Pledge *PledgeSession) EmergencyLendWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.EmergencyLendWithdrawal(&_Pledge.TransactOpts, _pid)
}

// EmergencyLendWithdrawal is a paid mutator transaction binding the contract method 0xbf38b8f6.
//
// Solidity: function emergencyLendWithdrawal(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) EmergencyLendWithdrawal(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.EmergencyLendWithdrawal(&_Pledge.TransactOpts, _pid)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) Finish(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "finish", _pid)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(uint256 _pid) returns()
func (_Pledge *PledgeSession) Finish(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Finish(&_Pledge.TransactOpts, _pid)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) Finish(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Finish(&_Pledge.TransactOpts, _pid)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) Liquidate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "liquidate", _pid)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _pid) returns()
func (_Pledge *PledgeSession) Liquidate(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Liquidate(&_Pledge.TransactOpts, _pid)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) Liquidate(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Liquidate(&_Pledge.TransactOpts, _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) RefundBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "refundBorrow", _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_Pledge *PledgeSession) RefundBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.RefundBorrow(&_Pledge.TransactOpts, _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) RefundBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.RefundBorrow(&_Pledge.TransactOpts, _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) RefundLend(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "refundLend", _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_Pledge *PledgeSession) RefundLend(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.RefundLend(&_Pledge.TransactOpts, _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) RefundLend(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.RefundLend(&_Pledge.TransactOpts, _pid)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_Pledge *PledgeTransactor) SetFee(opts *bind.TransactOpts, _lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setFee", _lendFee, _borrowFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_Pledge *PledgeSession) SetFee(_lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetFee(&_Pledge.TransactOpts, _lendFee, _borrowFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_Pledge *PledgeTransactorSession) SetFee(_lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetFee(&_Pledge.TransactOpts, _lendFee, _borrowFee)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Pledge *PledgeTransactor) SetFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setFeeAddress", _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Pledge *PledgeSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.SetFeeAddress(&_Pledge.TransactOpts, _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_Pledge *PledgeTransactorSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.SetFeeAddress(&_Pledge.TransactOpts, _feeAddress)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_Pledge *PledgeTransactor) SetMinAmount(opts *bind.TransactOpts, _minAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setMinAmount", _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_Pledge *PledgeSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetMinAmount(&_Pledge.TransactOpts, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_Pledge *PledgeTransactorSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetMinAmount(&_Pledge.TransactOpts, _minAmount)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Pledge *PledgeTransactor) SetPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setPause")
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Pledge *PledgeSession) SetPause() (*types.Transaction, error) {
	return _Pledge.Contract.SetPause(&_Pledge.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Pledge *PledgeTransactorSession) SetPause() (*types.Transaction, error) {
	return _Pledge.Contract.SetPause(&_Pledge.TransactOpts)
}

// SetSwapRouterAddress is a paid mutator transaction binding the contract method 0x5249961b.
//
// Solidity: function setSwapRouterAddress(address _swapRouter) returns()
func (_Pledge *PledgeTransactor) SetSwapRouterAddress(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setSwapRouterAddress", _swapRouter)
}

// SetSwapRouterAddress is a paid mutator transaction binding the contract method 0x5249961b.
//
// Solidity: function setSwapRouterAddress(address _swapRouter) returns()
func (_Pledge *PledgeSession) SetSwapRouterAddress(_swapRouter common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.SetSwapRouterAddress(&_Pledge.TransactOpts, _swapRouter)
}

// SetSwapRouterAddress is a paid mutator transaction binding the contract method 0x5249961b.
//
// Solidity: function setSwapRouterAddress(address _swapRouter) returns()
func (_Pledge *PledgeTransactorSession) SetSwapRouterAddress(_swapRouter common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.SetSwapRouterAddress(&_Pledge.TransactOpts, _swapRouter)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _pid) returns()
func (_Pledge *PledgeTransactor) Settle(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "settle", _pid)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _pid) returns()
func (_Pledge *PledgeSession) Settle(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Settle(&_Pledge.TransactOpts, _pid)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _pid) returns()
func (_Pledge *PledgeTransactorSession) Settle(_pid *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Settle(&_Pledge.TransactOpts, _pid)
}

// WithdrawBorrow is a paid mutator transaction binding the contract method 0x1e107979.
//
// Solidity: function withdrawBorrow(uint256 _pid, uint256 _jpAmount) returns()
func (_Pledge *PledgeTransactor) WithdrawBorrow(opts *bind.TransactOpts, _pid *big.Int, _jpAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "withdrawBorrow", _pid, _jpAmount)
}

// WithdrawBorrow is a paid mutator transaction binding the contract method 0x1e107979.
//
// Solidity: function withdrawBorrow(uint256 _pid, uint256 _jpAmount) returns()
func (_Pledge *PledgeSession) WithdrawBorrow(_pid *big.Int, _jpAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.WithdrawBorrow(&_Pledge.TransactOpts, _pid, _jpAmount)
}

// WithdrawBorrow is a paid mutator transaction binding the contract method 0x1e107979.
//
// Solidity: function withdrawBorrow(uint256 _pid, uint256 _jpAmount) returns()
func (_Pledge *PledgeTransactorSession) WithdrawBorrow(_pid *big.Int, _jpAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.WithdrawBorrow(&_Pledge.TransactOpts, _pid, _jpAmount)
}

// WithdrawLend is a paid mutator transaction binding the contract method 0x38f2aa76.
//
// Solidity: function withdrawLend(uint256 _pid, uint256 _spAmount) returns()
func (_Pledge *PledgeTransactor) WithdrawLend(opts *bind.TransactOpts, _pid *big.Int, _spAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "withdrawLend", _pid, _spAmount)
}

// WithdrawLend is a paid mutator transaction binding the contract method 0x38f2aa76.
//
// Solidity: function withdrawLend(uint256 _pid, uint256 _spAmount) returns()
func (_Pledge *PledgeSession) WithdrawLend(_pid *big.Int, _spAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.WithdrawLend(&_Pledge.TransactOpts, _pid, _spAmount)
}

// WithdrawLend is a paid mutator transaction binding the contract method 0x38f2aa76.
//
// Solidity: function withdrawLend(uint256 _pid, uint256 _spAmount) returns()
func (_Pledge *PledgeTransactorSession) WithdrawLend(_pid *big.Int, _spAmount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.WithdrawLend(&_Pledge.TransactOpts, _pid, _spAmount)
}

// PledgeClaimBorrowIterator is returned from FilterClaimBorrow and is used to iterate over the raw logs and unpacked data for ClaimBorrow events raised by the Pledge contract.
type PledgeClaimBorrowIterator struct {
	Event *PledgeClaimBorrow // Event containing the contract specifics and raw log

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
func (it *PledgeClaimBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeClaimBorrow)
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
		it.Event = new(PledgeClaimBorrow)
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
func (it *PledgeClaimBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeClaimBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeClaimBorrow represents a ClaimBorrow event raised by the Pledge contract.
type PledgeClaimBorrow struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimBorrow is a free log retrieval operation binding the contract event 0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd.
//
// Solidity: event ClaimBorrow(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) FilterClaimBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeClaimBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "ClaimBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeClaimBorrowIterator{contract: _Pledge.contract, event: "ClaimBorrow", logs: logs, sub: sub}, nil
}

// WatchClaimBorrow is a free log subscription operation binding the contract event 0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd.
//
// Solidity: event ClaimBorrow(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) WatchClaimBorrow(opts *bind.WatchOpts, sink chan<- *PledgeClaimBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "ClaimBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeClaimBorrow)
				if err := _Pledge.contract.UnpackLog(event, "ClaimBorrow", log); err != nil {
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

// ParseClaimBorrow is a log parse operation binding the contract event 0x3ddafe3ebb4d0c818317027aabfa82dc9983942ceeb80523167e2de047b17fbd.
//
// Solidity: event ClaimBorrow(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) ParseClaimBorrow(log types.Log) (*PledgeClaimBorrow, error) {
	event := new(PledgeClaimBorrow)
	if err := _Pledge.contract.UnpackLog(event, "ClaimBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeClaimLendIterator is returned from FilterClaimLend and is used to iterate over the raw logs and unpacked data for ClaimLend events raised by the Pledge contract.
type PledgeClaimLendIterator struct {
	Event *PledgeClaimLend // Event containing the contract specifics and raw log

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
func (it *PledgeClaimLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeClaimLend)
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
		it.Event = new(PledgeClaimLend)
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
func (it *PledgeClaimLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeClaimLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeClaimLend represents a ClaimLend event raised by the Pledge contract.
type PledgeClaimLend struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimLend is a free log retrieval operation binding the contract event 0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd.
//
// Solidity: event ClaimLend(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) FilterClaimLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeClaimLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "ClaimLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeClaimLendIterator{contract: _Pledge.contract, event: "ClaimLend", logs: logs, sub: sub}, nil
}

// WatchClaimLend is a free log subscription operation binding the contract event 0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd.
//
// Solidity: event ClaimLend(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) WatchClaimLend(opts *bind.WatchOpts, sink chan<- *PledgeClaimLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "ClaimLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeClaimLend)
				if err := _Pledge.contract.UnpackLog(event, "ClaimLend", log); err != nil {
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

// ParseClaimLend is a log parse operation binding the contract event 0x6f4dd2687b3c3bfa99d39742b01d6e0ad9604c48559791d5df4ff5df44b41dfd.
//
// Solidity: event ClaimLend(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) ParseClaimLend(log types.Log) (*PledgeClaimLend, error) {
	event := new(PledgeClaimLend)
	if err := _Pledge.contract.UnpackLog(event, "ClaimLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeDepositBorrowIterator is returned from FilterDepositBorrow and is used to iterate over the raw logs and unpacked data for DepositBorrow events raised by the Pledge contract.
type PledgeDepositBorrowIterator struct {
	Event *PledgeDepositBorrow // Event containing the contract specifics and raw log

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
func (it *PledgeDepositBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeDepositBorrow)
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
		it.Event = new(PledgeDepositBorrow)
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
func (it *PledgeDepositBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeDepositBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeDepositBorrow represents a DepositBorrow event raised by the Pledge contract.
type PledgeDepositBorrow struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositBorrow is a free log retrieval operation binding the contract event 0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5.
//
// Solidity: event DepositBorrow(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_Pledge *PledgeFilterer) FilterDepositBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeDepositBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "DepositBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeDepositBorrowIterator{contract: _Pledge.contract, event: "DepositBorrow", logs: logs, sub: sub}, nil
}

// WatchDepositBorrow is a free log subscription operation binding the contract event 0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5.
//
// Solidity: event DepositBorrow(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_Pledge *PledgeFilterer) WatchDepositBorrow(opts *bind.WatchOpts, sink chan<- *PledgeDepositBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "DepositBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeDepositBorrow)
				if err := _Pledge.contract.UnpackLog(event, "DepositBorrow", log); err != nil {
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

// ParseDepositBorrow is a log parse operation binding the contract event 0x1d7b72e666a0b6217efe7cfa1b604ea5c7b39219563ce48b30c9da77045247a5.
//
// Solidity: event DepositBorrow(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_Pledge *PledgeFilterer) ParseDepositBorrow(log types.Log) (*PledgeDepositBorrow, error) {
	event := new(PledgeDepositBorrow)
	if err := _Pledge.contract.UnpackLog(event, "DepositBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeDepositLendIterator is returned from FilterDepositLend and is used to iterate over the raw logs and unpacked data for DepositLend events raised by the Pledge contract.
type PledgeDepositLendIterator struct {
	Event *PledgeDepositLend // Event containing the contract specifics and raw log

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
func (it *PledgeDepositLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeDepositLend)
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
		it.Event = new(PledgeDepositLend)
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
func (it *PledgeDepositLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeDepositLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeDepositLend represents a DepositLend event raised by the Pledge contract.
type PledgeDepositLend struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositLend is a free log retrieval operation binding the contract event 0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea.
//
// Solidity: event DepositLend(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_Pledge *PledgeFilterer) FilterDepositLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeDepositLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "DepositLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeDepositLendIterator{contract: _Pledge.contract, event: "DepositLend", logs: logs, sub: sub}, nil
}

// WatchDepositLend is a free log subscription operation binding the contract event 0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea.
//
// Solidity: event DepositLend(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_Pledge *PledgeFilterer) WatchDepositLend(opts *bind.WatchOpts, sink chan<- *PledgeDepositLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "DepositLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeDepositLend)
				if err := _Pledge.contract.UnpackLog(event, "DepositLend", log); err != nil {
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

// ParseDepositLend is a log parse operation binding the contract event 0x129e8c18c2f7baf99c7eb257934c21f038c72412803512dcf0a942a4562a82ea.
//
// Solidity: event DepositLend(address indexed from, address indexed token, uint256 amount, uint256 mintAmount)
func (_Pledge *PledgeFilterer) ParseDepositLend(log types.Log) (*PledgeDepositLend, error) {
	event := new(PledgeDepositLend)
	if err := _Pledge.contract.UnpackLog(event, "DepositLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeEmergencyBorrowWithdrawalIterator is returned from FilterEmergencyBorrowWithdrawal and is used to iterate over the raw logs and unpacked data for EmergencyBorrowWithdrawal events raised by the Pledge contract.
type PledgeEmergencyBorrowWithdrawalIterator struct {
	Event *PledgeEmergencyBorrowWithdrawal // Event containing the contract specifics and raw log

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
func (it *PledgeEmergencyBorrowWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeEmergencyBorrowWithdrawal)
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
		it.Event = new(PledgeEmergencyBorrowWithdrawal)
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
func (it *PledgeEmergencyBorrowWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeEmergencyBorrowWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeEmergencyBorrowWithdrawal represents a EmergencyBorrowWithdrawal event raised by the Pledge contract.
type PledgeEmergencyBorrowWithdrawal struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyBorrowWithdrawal is a free log retrieval operation binding the contract event 0x5a06c7de92f1dc59e8cba872927d016c80ce5f0fb2295c898dfb7a2f08e43fb1.
//
// Solidity: event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) FilterEmergencyBorrowWithdrawal(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeEmergencyBorrowWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "EmergencyBorrowWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeEmergencyBorrowWithdrawalIterator{contract: _Pledge.contract, event: "EmergencyBorrowWithdrawal", logs: logs, sub: sub}, nil
}

// WatchEmergencyBorrowWithdrawal is a free log subscription operation binding the contract event 0x5a06c7de92f1dc59e8cba872927d016c80ce5f0fb2295c898dfb7a2f08e43fb1.
//
// Solidity: event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) WatchEmergencyBorrowWithdrawal(opts *bind.WatchOpts, sink chan<- *PledgeEmergencyBorrowWithdrawal, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "EmergencyBorrowWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeEmergencyBorrowWithdrawal)
				if err := _Pledge.contract.UnpackLog(event, "EmergencyBorrowWithdrawal", log); err != nil {
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

// ParseEmergencyBorrowWithdrawal is a log parse operation binding the contract event 0x5a06c7de92f1dc59e8cba872927d016c80ce5f0fb2295c898dfb7a2f08e43fb1.
//
// Solidity: event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) ParseEmergencyBorrowWithdrawal(log types.Log) (*PledgeEmergencyBorrowWithdrawal, error) {
	event := new(PledgeEmergencyBorrowWithdrawal)
	if err := _Pledge.contract.UnpackLog(event, "EmergencyBorrowWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeEmergencyLendWithdrawalIterator is returned from FilterEmergencyLendWithdrawal and is used to iterate over the raw logs and unpacked data for EmergencyLendWithdrawal events raised by the Pledge contract.
type PledgeEmergencyLendWithdrawalIterator struct {
	Event *PledgeEmergencyLendWithdrawal // Event containing the contract specifics and raw log

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
func (it *PledgeEmergencyLendWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeEmergencyLendWithdrawal)
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
		it.Event = new(PledgeEmergencyLendWithdrawal)
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
func (it *PledgeEmergencyLendWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeEmergencyLendWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeEmergencyLendWithdrawal represents a EmergencyLendWithdrawal event raised by the Pledge contract.
type PledgeEmergencyLendWithdrawal struct {
	From   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyLendWithdrawal is a free log retrieval operation binding the contract event 0x71d14c5f08cb34cbfb59c06ea5151aafbf742d0b6ed00fdb83addd9afb5c0fd0.
//
// Solidity: event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) FilterEmergencyLendWithdrawal(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeEmergencyLendWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "EmergencyLendWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeEmergencyLendWithdrawalIterator{contract: _Pledge.contract, event: "EmergencyLendWithdrawal", logs: logs, sub: sub}, nil
}

// WatchEmergencyLendWithdrawal is a free log subscription operation binding the contract event 0x71d14c5f08cb34cbfb59c06ea5151aafbf742d0b6ed00fdb83addd9afb5c0fd0.
//
// Solidity: event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) WatchEmergencyLendWithdrawal(opts *bind.WatchOpts, sink chan<- *PledgeEmergencyLendWithdrawal, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "EmergencyLendWithdrawal", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeEmergencyLendWithdrawal)
				if err := _Pledge.contract.UnpackLog(event, "EmergencyLendWithdrawal", log); err != nil {
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

// ParseEmergencyLendWithdrawal is a log parse operation binding the contract event 0x71d14c5f08cb34cbfb59c06ea5151aafbf742d0b6ed00fdb83addd9afb5c0fd0.
//
// Solidity: event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) ParseEmergencyLendWithdrawal(log types.Log) (*PledgeEmergencyLendWithdrawal, error) {
	event := new(PledgeEmergencyLendWithdrawal)
	if err := _Pledge.contract.UnpackLog(event, "EmergencyLendWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Pledge contract.
type PledgeRedeemIterator struct {
	Event *PledgeRedeem // Event containing the contract specifics and raw log

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
func (it *PledgeRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeRedeem)
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
		it.Event = new(PledgeRedeem)
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
func (it *PledgeRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeRedeem represents a Redeem event raised by the Pledge contract.
type PledgeRedeem struct {
	Recieptor common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed recieptor, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) FilterRedeem(opts *bind.FilterOpts, recieptor []common.Address, token []common.Address) (*PledgeRedeemIterator, error) {

	var recieptorRule []interface{}
	for _, recieptorItem := range recieptor {
		recieptorRule = append(recieptorRule, recieptorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "Redeem", recieptorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeRedeemIterator{contract: _Pledge.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed recieptor, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *PledgeRedeem, recieptor []common.Address, token []common.Address) (event.Subscription, error) {

	var recieptorRule []interface{}
	for _, recieptorItem := range recieptor {
		recieptorRule = append(recieptorRule, recieptorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "Redeem", recieptorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeRedeem)
				if err := _Pledge.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed recieptor, address indexed token, uint256 amount)
func (_Pledge *PledgeFilterer) ParseRedeem(log types.Log) (*PledgeRedeem, error) {
	event := new(PledgeRedeem)
	if err := _Pledge.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeRefundBorrowIterator is returned from FilterRefundBorrow and is used to iterate over the raw logs and unpacked data for RefundBorrow events raised by the Pledge contract.
type PledgeRefundBorrowIterator struct {
	Event *PledgeRefundBorrow // Event containing the contract specifics and raw log

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
func (it *PledgeRefundBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeRefundBorrow)
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
		it.Event = new(PledgeRefundBorrow)
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
func (it *PledgeRefundBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeRefundBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeRefundBorrow represents a RefundBorrow event raised by the Pledge contract.
type PledgeRefundBorrow struct {
	From   common.Address
	Token  common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundBorrow is a free log retrieval operation binding the contract event 0x732816f48de550f238bd0d4f5b79819c7b24a49d6132928978e3cd36568dd5db.
//
// Solidity: event RefundBorrow(address indexed from, address indexed token, uint256 refund)
func (_Pledge *PledgeFilterer) FilterRefundBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeRefundBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "RefundBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeRefundBorrowIterator{contract: _Pledge.contract, event: "RefundBorrow", logs: logs, sub: sub}, nil
}

// WatchRefundBorrow is a free log subscription operation binding the contract event 0x732816f48de550f238bd0d4f5b79819c7b24a49d6132928978e3cd36568dd5db.
//
// Solidity: event RefundBorrow(address indexed from, address indexed token, uint256 refund)
func (_Pledge *PledgeFilterer) WatchRefundBorrow(opts *bind.WatchOpts, sink chan<- *PledgeRefundBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "RefundBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeRefundBorrow)
				if err := _Pledge.contract.UnpackLog(event, "RefundBorrow", log); err != nil {
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

// ParseRefundBorrow is a log parse operation binding the contract event 0x732816f48de550f238bd0d4f5b79819c7b24a49d6132928978e3cd36568dd5db.
//
// Solidity: event RefundBorrow(address indexed from, address indexed token, uint256 refund)
func (_Pledge *PledgeFilterer) ParseRefundBorrow(log types.Log) (*PledgeRefundBorrow, error) {
	event := new(PledgeRefundBorrow)
	if err := _Pledge.contract.UnpackLog(event, "RefundBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeRefundLendIterator is returned from FilterRefundLend and is used to iterate over the raw logs and unpacked data for RefundLend events raised by the Pledge contract.
type PledgeRefundLendIterator struct {
	Event *PledgeRefundLend // Event containing the contract specifics and raw log

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
func (it *PledgeRefundLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeRefundLend)
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
		it.Event = new(PledgeRefundLend)
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
func (it *PledgeRefundLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeRefundLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeRefundLend represents a RefundLend event raised by the Pledge contract.
type PledgeRefundLend struct {
	From   common.Address
	Token  common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundLend is a free log retrieval operation binding the contract event 0xc3e20279d41b3ed21d277920877e5e5c6665bf6aca607046a3fe0fd2bd6bda7d.
//
// Solidity: event RefundLend(address indexed from, address indexed token, uint256 refund)
func (_Pledge *PledgeFilterer) FilterRefundLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeRefundLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "RefundLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeRefundLendIterator{contract: _Pledge.contract, event: "RefundLend", logs: logs, sub: sub}, nil
}

// WatchRefundLend is a free log subscription operation binding the contract event 0xc3e20279d41b3ed21d277920877e5e5c6665bf6aca607046a3fe0fd2bd6bda7d.
//
// Solidity: event RefundLend(address indexed from, address indexed token, uint256 refund)
func (_Pledge *PledgeFilterer) WatchRefundLend(opts *bind.WatchOpts, sink chan<- *PledgeRefundLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "RefundLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeRefundLend)
				if err := _Pledge.contract.UnpackLog(event, "RefundLend", log); err != nil {
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

// ParseRefundLend is a log parse operation binding the contract event 0xc3e20279d41b3ed21d277920877e5e5c6665bf6aca607046a3fe0fd2bd6bda7d.
//
// Solidity: event RefundLend(address indexed from, address indexed token, uint256 refund)
func (_Pledge *PledgeFilterer) ParseRefundLend(log types.Log) (*PledgeRefundLend, error) {
	event := new(PledgeRefundLend)
	if err := _Pledge.contract.UnpackLog(event, "RefundLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Pledge contract.
type PledgeSetFeeIterator struct {
	Event *PledgeSetFee // Event containing the contract specifics and raw log

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
func (it *PledgeSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeSetFee)
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
		it.Event = new(PledgeSetFee)
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
func (it *PledgeSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeSetFee represents a SetFee event raised by the Pledge contract.
type PledgeSetFee struct {
	NewLendFee   *big.Int
	NewBorrowFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_Pledge *PledgeFilterer) FilterSetFee(opts *bind.FilterOpts, newLendFee []*big.Int, newBorrowFee []*big.Int) (*PledgeSetFeeIterator, error) {

	var newLendFeeRule []interface{}
	for _, newLendFeeItem := range newLendFee {
		newLendFeeRule = append(newLendFeeRule, newLendFeeItem)
	}
	var newBorrowFeeRule []interface{}
	for _, newBorrowFeeItem := range newBorrowFee {
		newBorrowFeeRule = append(newBorrowFeeRule, newBorrowFeeItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "SetFee", newLendFeeRule, newBorrowFeeRule)
	if err != nil {
		return nil, err
	}
	return &PledgeSetFeeIterator{contract: _Pledge.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_Pledge *PledgeFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *PledgeSetFee, newLendFee []*big.Int, newBorrowFee []*big.Int) (event.Subscription, error) {

	var newLendFeeRule []interface{}
	for _, newLendFeeItem := range newLendFee {
		newLendFeeRule = append(newLendFeeRule, newLendFeeItem)
	}
	var newBorrowFeeRule []interface{}
	for _, newBorrowFeeItem := range newBorrowFee {
		newBorrowFeeRule = append(newBorrowFeeRule, newBorrowFeeItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "SetFee", newLendFeeRule, newBorrowFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeSetFee)
				if err := _Pledge.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_Pledge *PledgeFilterer) ParseSetFee(log types.Log) (*PledgeSetFee, error) {
	event := new(PledgeSetFee)
	if err := _Pledge.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeSetFeeAddressIterator is returned from FilterSetFeeAddress and is used to iterate over the raw logs and unpacked data for SetFeeAddress events raised by the Pledge contract.
type PledgeSetFeeAddressIterator struct {
	Event *PledgeSetFeeAddress // Event containing the contract specifics and raw log

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
func (it *PledgeSetFeeAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeSetFeeAddress)
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
		it.Event = new(PledgeSetFeeAddress)
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
func (it *PledgeSetFeeAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeSetFeeAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeSetFeeAddress represents a SetFeeAddress event raised by the Pledge contract.
type PledgeSetFeeAddress struct {
	OldFeeAddress common.Address
	NewFeeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAddress is a free log retrieval operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_Pledge *PledgeFilterer) FilterSetFeeAddress(opts *bind.FilterOpts, oldFeeAddress []common.Address, newFeeAddress []common.Address) (*PledgeSetFeeAddressIterator, error) {

	var oldFeeAddressRule []interface{}
	for _, oldFeeAddressItem := range oldFeeAddress {
		oldFeeAddressRule = append(oldFeeAddressRule, oldFeeAddressItem)
	}
	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "SetFeeAddress", oldFeeAddressRule, newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return &PledgeSetFeeAddressIterator{contract: _Pledge.contract, event: "SetFeeAddress", logs: logs, sub: sub}, nil
}

// WatchSetFeeAddress is a free log subscription operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_Pledge *PledgeFilterer) WatchSetFeeAddress(opts *bind.WatchOpts, sink chan<- *PledgeSetFeeAddress, oldFeeAddress []common.Address, newFeeAddress []common.Address) (event.Subscription, error) {

	var oldFeeAddressRule []interface{}
	for _, oldFeeAddressItem := range oldFeeAddress {
		oldFeeAddressRule = append(oldFeeAddressRule, oldFeeAddressItem)
	}
	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "SetFeeAddress", oldFeeAddressRule, newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeSetFeeAddress)
				if err := _Pledge.contract.UnpackLog(event, "SetFeeAddress", log); err != nil {
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

// ParseSetFeeAddress is a log parse operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_Pledge *PledgeFilterer) ParseSetFeeAddress(log types.Log) (*PledgeSetFeeAddress, error) {
	event := new(PledgeSetFeeAddress)
	if err := _Pledge.contract.UnpackLog(event, "SetFeeAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeSetMinAmountIterator is returned from FilterSetMinAmount and is used to iterate over the raw logs and unpacked data for SetMinAmount events raised by the Pledge contract.
type PledgeSetMinAmountIterator struct {
	Event *PledgeSetMinAmount // Event containing the contract specifics and raw log

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
func (it *PledgeSetMinAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeSetMinAmount)
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
		it.Event = new(PledgeSetMinAmount)
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
func (it *PledgeSetMinAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeSetMinAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeSetMinAmount represents a SetMinAmount event raised by the Pledge contract.
type PledgeSetMinAmount struct {
	OldMinAmount *big.Int
	NewMinAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetMinAmount is a free log retrieval operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_Pledge *PledgeFilterer) FilterSetMinAmount(opts *bind.FilterOpts, oldMinAmount []*big.Int, newMinAmount []*big.Int) (*PledgeSetMinAmountIterator, error) {

	var oldMinAmountRule []interface{}
	for _, oldMinAmountItem := range oldMinAmount {
		oldMinAmountRule = append(oldMinAmountRule, oldMinAmountItem)
	}
	var newMinAmountRule []interface{}
	for _, newMinAmountItem := range newMinAmount {
		newMinAmountRule = append(newMinAmountRule, newMinAmountItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "SetMinAmount", oldMinAmountRule, newMinAmountRule)
	if err != nil {
		return nil, err
	}
	return &PledgeSetMinAmountIterator{contract: _Pledge.contract, event: "SetMinAmount", logs: logs, sub: sub}, nil
}

// WatchSetMinAmount is a free log subscription operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_Pledge *PledgeFilterer) WatchSetMinAmount(opts *bind.WatchOpts, sink chan<- *PledgeSetMinAmount, oldMinAmount []*big.Int, newMinAmount []*big.Int) (event.Subscription, error) {

	var oldMinAmountRule []interface{}
	for _, oldMinAmountItem := range oldMinAmount {
		oldMinAmountRule = append(oldMinAmountRule, oldMinAmountItem)
	}
	var newMinAmountRule []interface{}
	for _, newMinAmountItem := range newMinAmount {
		newMinAmountRule = append(newMinAmountRule, newMinAmountItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "SetMinAmount", oldMinAmountRule, newMinAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeSetMinAmount)
				if err := _Pledge.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
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

// ParseSetMinAmount is a log parse operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_Pledge *PledgeFilterer) ParseSetMinAmount(log types.Log) (*PledgeSetMinAmount, error) {
	event := new(PledgeSetMinAmount)
	if err := _Pledge.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeSetSwapRouterAddressIterator is returned from FilterSetSwapRouterAddress and is used to iterate over the raw logs and unpacked data for SetSwapRouterAddress events raised by the Pledge contract.
type PledgeSetSwapRouterAddressIterator struct {
	Event *PledgeSetSwapRouterAddress // Event containing the contract specifics and raw log

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
func (it *PledgeSetSwapRouterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeSetSwapRouterAddress)
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
		it.Event = new(PledgeSetSwapRouterAddress)
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
func (it *PledgeSetSwapRouterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeSetSwapRouterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeSetSwapRouterAddress represents a SetSwapRouterAddress event raised by the Pledge contract.
type PledgeSetSwapRouterAddress struct {
	OldSwapAddress common.Address
	NewSwapAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSwapRouterAddress is a free log retrieval operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_Pledge *PledgeFilterer) FilterSetSwapRouterAddress(opts *bind.FilterOpts, oldSwapAddress []common.Address, newSwapAddress []common.Address) (*PledgeSetSwapRouterAddressIterator, error) {

	var oldSwapAddressRule []interface{}
	for _, oldSwapAddressItem := range oldSwapAddress {
		oldSwapAddressRule = append(oldSwapAddressRule, oldSwapAddressItem)
	}
	var newSwapAddressRule []interface{}
	for _, newSwapAddressItem := range newSwapAddress {
		newSwapAddressRule = append(newSwapAddressRule, newSwapAddressItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "SetSwapRouterAddress", oldSwapAddressRule, newSwapAddressRule)
	if err != nil {
		return nil, err
	}
	return &PledgeSetSwapRouterAddressIterator{contract: _Pledge.contract, event: "SetSwapRouterAddress", logs: logs, sub: sub}, nil
}

// WatchSetSwapRouterAddress is a free log subscription operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_Pledge *PledgeFilterer) WatchSetSwapRouterAddress(opts *bind.WatchOpts, sink chan<- *PledgeSetSwapRouterAddress, oldSwapAddress []common.Address, newSwapAddress []common.Address) (event.Subscription, error) {

	var oldSwapAddressRule []interface{}
	for _, oldSwapAddressItem := range oldSwapAddress {
		oldSwapAddressRule = append(oldSwapAddressRule, oldSwapAddressItem)
	}
	var newSwapAddressRule []interface{}
	for _, newSwapAddressItem := range newSwapAddress {
		newSwapAddressRule = append(newSwapAddressRule, newSwapAddressItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "SetSwapRouterAddress", oldSwapAddressRule, newSwapAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeSetSwapRouterAddress)
				if err := _Pledge.contract.UnpackLog(event, "SetSwapRouterAddress", log); err != nil {
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

// ParseSetSwapRouterAddress is a log parse operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_Pledge *PledgeFilterer) ParseSetSwapRouterAddress(log types.Log) (*PledgeSetSwapRouterAddress, error) {
	event := new(PledgeSetSwapRouterAddress)
	if err := _Pledge.contract.UnpackLog(event, "SetSwapRouterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeStateChangeIterator is returned from FilterStateChange and is used to iterate over the raw logs and unpacked data for StateChange events raised by the Pledge contract.
type PledgeStateChangeIterator struct {
	Event *PledgeStateChange // Event containing the contract specifics and raw log

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
func (it *PledgeStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeStateChange)
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
		it.Event = new(PledgeStateChange)
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
func (it *PledgeStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeStateChange represents a StateChange event raised by the Pledge contract.
type PledgeStateChange struct {
	Pid         *big.Int
	BeforeState *big.Int
	AfterState  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStateChange is a free log retrieval operation binding the contract event 0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb.
//
// Solidity: event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState)
func (_Pledge *PledgeFilterer) FilterStateChange(opts *bind.FilterOpts, pid []*big.Int, beforeState []*big.Int, afterState []*big.Int) (*PledgeStateChangeIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var beforeStateRule []interface{}
	for _, beforeStateItem := range beforeState {
		beforeStateRule = append(beforeStateRule, beforeStateItem)
	}
	var afterStateRule []interface{}
	for _, afterStateItem := range afterState {
		afterStateRule = append(afterStateRule, afterStateItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "StateChange", pidRule, beforeStateRule, afterStateRule)
	if err != nil {
		return nil, err
	}
	return &PledgeStateChangeIterator{contract: _Pledge.contract, event: "StateChange", logs: logs, sub: sub}, nil
}

// WatchStateChange is a free log subscription operation binding the contract event 0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb.
//
// Solidity: event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState)
func (_Pledge *PledgeFilterer) WatchStateChange(opts *bind.WatchOpts, sink chan<- *PledgeStateChange, pid []*big.Int, beforeState []*big.Int, afterState []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var beforeStateRule []interface{}
	for _, beforeStateItem := range beforeState {
		beforeStateRule = append(beforeStateRule, beforeStateItem)
	}
	var afterStateRule []interface{}
	for _, afterStateItem := range afterState {
		afterStateRule = append(afterStateRule, afterStateItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "StateChange", pidRule, beforeStateRule, afterStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeStateChange)
				if err := _Pledge.contract.UnpackLog(event, "StateChange", log); err != nil {
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

// ParseStateChange is a log parse operation binding the contract event 0x516112f3bf06e373fcea44db364769c04cc7ef4392e6de95d2b250720bcacefb.
//
// Solidity: event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState)
func (_Pledge *PledgeFilterer) ParseStateChange(log types.Log) (*PledgeStateChange, error) {
	event := new(PledgeStateChange)
	if err := _Pledge.contract.UnpackLog(event, "StateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pledge contract.
type PledgeSwapIterator struct {
	Event *PledgeSwap // Event containing the contract specifics and raw log

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
func (it *PledgeSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeSwap)
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
		it.Event = new(PledgeSwap)
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
func (it *PledgeSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeSwap represents a Swap event raised by the Pledge contract.
type PledgeSwap struct {
	FromCoin  common.Address
	ToCoin    common.Address
	FromValue *big.Int
	ToValue   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xfa2dda1cc1b86e41239702756b13effbc1a092b5c57e3ad320fbe4f3b13fe235.
//
// Solidity: event Swap(address indexed fromCoin, address indexed toCoin, uint256 fromValue, uint256 toValue)
func (_Pledge *PledgeFilterer) FilterSwap(opts *bind.FilterOpts, fromCoin []common.Address, toCoin []common.Address) (*PledgeSwapIterator, error) {

	var fromCoinRule []interface{}
	for _, fromCoinItem := range fromCoin {
		fromCoinRule = append(fromCoinRule, fromCoinItem)
	}
	var toCoinRule []interface{}
	for _, toCoinItem := range toCoin {
		toCoinRule = append(toCoinRule, toCoinItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "Swap", fromCoinRule, toCoinRule)
	if err != nil {
		return nil, err
	}
	return &PledgeSwapIterator{contract: _Pledge.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xfa2dda1cc1b86e41239702756b13effbc1a092b5c57e3ad320fbe4f3b13fe235.
//
// Solidity: event Swap(address indexed fromCoin, address indexed toCoin, uint256 fromValue, uint256 toValue)
func (_Pledge *PledgeFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PledgeSwap, fromCoin []common.Address, toCoin []common.Address) (event.Subscription, error) {

	var fromCoinRule []interface{}
	for _, fromCoinItem := range fromCoin {
		fromCoinRule = append(fromCoinRule, fromCoinItem)
	}
	var toCoinRule []interface{}
	for _, toCoinItem := range toCoin {
		toCoinRule = append(toCoinRule, toCoinItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "Swap", fromCoinRule, toCoinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeSwap)
				if err := _Pledge.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xfa2dda1cc1b86e41239702756b13effbc1a092b5c57e3ad320fbe4f3b13fe235.
//
// Solidity: event Swap(address indexed fromCoin, address indexed toCoin, uint256 fromValue, uint256 toValue)
func (_Pledge *PledgeFilterer) ParseSwap(log types.Log) (*PledgeSwap, error) {
	event := new(PledgeSwap)
	if err := _Pledge.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeWithdrawBorrowIterator is returned from FilterWithdrawBorrow and is used to iterate over the raw logs and unpacked data for WithdrawBorrow events raised by the Pledge contract.
type PledgeWithdrawBorrowIterator struct {
	Event *PledgeWithdrawBorrow // Event containing the contract specifics and raw log

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
func (it *PledgeWithdrawBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeWithdrawBorrow)
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
		it.Event = new(PledgeWithdrawBorrow)
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
func (it *PledgeWithdrawBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeWithdrawBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeWithdrawBorrow represents a WithdrawBorrow event raised by the Pledge contract.
type PledgeWithdrawBorrow struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	BurnAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBorrow is a free log retrieval operation binding the contract event 0x0f5e74952c2f9259a748f3aa9a6c4534a6f46a5966e5baabdb6bd337f05234a8.
//
// Solidity: event WithdrawBorrow(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_Pledge *PledgeFilterer) FilterWithdrawBorrow(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeWithdrawBorrowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "WithdrawBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeWithdrawBorrowIterator{contract: _Pledge.contract, event: "WithdrawBorrow", logs: logs, sub: sub}, nil
}

// WatchWithdrawBorrow is a free log subscription operation binding the contract event 0x0f5e74952c2f9259a748f3aa9a6c4534a6f46a5966e5baabdb6bd337f05234a8.
//
// Solidity: event WithdrawBorrow(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_Pledge *PledgeFilterer) WatchWithdrawBorrow(opts *bind.WatchOpts, sink chan<- *PledgeWithdrawBorrow, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "WithdrawBorrow", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeWithdrawBorrow)
				if err := _Pledge.contract.UnpackLog(event, "WithdrawBorrow", log); err != nil {
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

// ParseWithdrawBorrow is a log parse operation binding the contract event 0x0f5e74952c2f9259a748f3aa9a6c4534a6f46a5966e5baabdb6bd337f05234a8.
//
// Solidity: event WithdrawBorrow(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_Pledge *PledgeFilterer) ParseWithdrawBorrow(log types.Log) (*PledgeWithdrawBorrow, error) {
	event := new(PledgeWithdrawBorrow)
	if err := _Pledge.contract.UnpackLog(event, "WithdrawBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeWithdrawLendIterator is returned from FilterWithdrawLend and is used to iterate over the raw logs and unpacked data for WithdrawLend events raised by the Pledge contract.
type PledgeWithdrawLendIterator struct {
	Event *PledgeWithdrawLend // Event containing the contract specifics and raw log

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
func (it *PledgeWithdrawLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeWithdrawLend)
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
		it.Event = new(PledgeWithdrawLend)
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
func (it *PledgeWithdrawLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeWithdrawLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeWithdrawLend represents a WithdrawLend event raised by the Pledge contract.
type PledgeWithdrawLend struct {
	From       common.Address
	Token      common.Address
	Amount     *big.Int
	BurnAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLend is a free log retrieval operation binding the contract event 0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d.
//
// Solidity: event WithdrawLend(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_Pledge *PledgeFilterer) FilterWithdrawLend(opts *bind.FilterOpts, from []common.Address, token []common.Address) (*PledgeWithdrawLendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "WithdrawLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PledgeWithdrawLendIterator{contract: _Pledge.contract, event: "WithdrawLend", logs: logs, sub: sub}, nil
}

// WatchWithdrawLend is a free log subscription operation binding the contract event 0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d.
//
// Solidity: event WithdrawLend(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_Pledge *PledgeFilterer) WatchWithdrawLend(opts *bind.WatchOpts, sink chan<- *PledgeWithdrawLend, from []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "WithdrawLend", fromRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeWithdrawLend)
				if err := _Pledge.contract.UnpackLog(event, "WithdrawLend", log); err != nil {
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

// ParseWithdrawLend is a log parse operation binding the contract event 0x690f32ccf3e832d5ff975d781039bc2affebee9c973939c9b710091b87954c9d.
//
// Solidity: event WithdrawLend(address indexed from, address indexed token, uint256 amount, uint256 burnAmount)
func (_Pledge *PledgeFilterer) ParseWithdrawLend(log types.Log) (*PledgeWithdrawLend, error) {
	event := new(PledgeWithdrawLend)
	if err := _Pledge.contract.UnpackLog(event, "WithdrawLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
