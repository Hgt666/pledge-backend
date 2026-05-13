package parser

import (
	"easy-swap/logger"
	"easy-swap/utils"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

var (
	pledgeAbi abi.ABI
)

// Init 初始化abi解析器
func Init() error {
	// 读取 abi文件
	abiBytes, err := utils.ReadABIFile("PledgePool.abi")
	if err != nil {
		return err
	}
	// 解析 abi
	parsedAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return err
	}
	pledgeAbi = parsedAbi
	return nil
}

// 事件
// // 存款借出事件，from是借出者地址，token是借出的代币地址，amount是借出的数量，mintAmount是生成的数量

// // 借出退款事件，from是退款者地址，token是退款的代币地址，refund是退款的数量
// event RefundLend(address indexed from, address indexed token, uint256 refund); 
// // 借出索赔事件，from是索赔者地址，token是索赔的代币地址，amount是索赔的数量
// event ClaimLend(address indexed from, address indexed token, uint256 amount); 
// // 提取借出事件，from是提取者地址，token是提取的代币地址，amount是提取的数量，burnAmount是销毁的数量
// event WithdrawLend(address indexed from,address indexed token,uint256 amount,uint256 burnAmount);

// // 借入退款事件，from是退款者地址，token是退款的代币地址，refund是退款的数量
// event RefundBorrow(address indexed from, address indexed token, uint256 refund);
// // 借入索赔事件，from是索赔者地址，token是索赔的代币地址，amount是索赔的数量
// event ClaimBorrow(address indexed from, address indexed token, uint256 amount); 
// // 提取借入事件，from是提取者地址，token是提取的代币地址，amount是提取的数量，burnAmount是销毁的数量
// event WithdrawBorrow(address indexed from,address indexed token,uint256 amount,uint256 burnAmount); 
// // 交换事件，fromCoin是交换前的币种地址，toCoin是交换后的币种地址，fromValue是交换前的数量，toValue是交换后的数量
// event Swap(address indexed fromCoin,address indexed toCoin,uint256 fromValue,uint256 toValue); 
// // 紧急借入提取事件，from是提取者地址，token是提取的代币地址，amount是提取的数量
// event EmergencyBorrowWithdrawal(address indexed from, address indexed token, uint256 amount); 
// // 紧急借出提取事件，from是提取者地址，token是提取的代币地址，amount是提取的数量
// event EmergencyLendWithdrawal(address indexed from, address indexed token, uint256 amount);
// // 状态改变事件，pid是项目id，beforeState是改变前的状态，afterState是改变后的状态
// event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState); 
// // 设置费用事件，newLendFee是新的借出费用，newBorrowFee是新的借入费用
// event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee);
// // 设置交换路由器地址事件，oldSwapAddress是旧的交换地址，newSwapAddress是新的交换地址
// event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress); 
// // 设置费用地址事件，oldFeeAddress是旧的费用地址，newFeeAddress是新的费用地址
// event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress);
// // 设置最小数量事件，oldMinAmount是旧的最小数量，newMinAmount是新的最小数量
// event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount); 

// event DepositLend(address indexed from,address indexed token,uint256 amount,uint256 mintAmount); 
// // 存款借入事件，from是借入者地址，token是借入的代币地址，amount是借入的数量，mintAmount是生成的数量
// event DepositBorrow(address indexed from,address indexed token,uint256 amount,uint256 mintAmount); 

type DepositLend struct {
	From    common.Address
	Token   common.Address
	Amount  *big.Int
	MintAmount *big.Int
}


type DepositBorrow struct {
	From    common.Address
	Token   common.Address
	Amount  *big.Int
	MintAmount *big.Int
}

// // 状态改变事件，pid是项目id，beforeState是改变前的状态，afterState是改变后的状态
// event StateChange(uint256 indexed pid, uint256 indexed beforeState, uint256 indexed afterState); 
type StateChange struct {
	Pid         *big.Int
	BeforeState *big.Int
	AfterState  *big.Int
}


// ParseDepositLend 解析 DepositLend 事件日志
func ParseDepositLend(log *types.Log) (*DepositLend, error) {
	var event DepositLend
	if err := pledgeAbi.UnpackIntoInterface(&event, "DepositLend", log.Data); err != nil {
		return nil, err
	}
	event.From = common.BytesToAddress(log.Topics[1].Bytes())
	event.Token = common.BytesToAddress(log.Topics[2].Bytes())
	logger.Log.Info("解析 DepositLend事件成功",zap.String("from",event.From.String()))
	return &event, nil
}

// ParseDepositBorrow 解析 DepositBorrow 事件日志
func ParseDepositBorrow(log *types.Log) (*DepositBorrow, error) {
	var event DepositBorrow
	if err := pledgeAbi.UnpackIntoInterface(&event, "DepositBorrow", log.Data); err != nil {
		return nil, err
	}
	event.From = common.BytesToAddress(log.Topics[1].Bytes())
	event.Token = common.BytesToAddress(log.Topics[2].Bytes())
	logger.Log.Info("解析 DepositBorrow事件成功",zap.String("from",event.From.String()))
	return &event, nil
}

// ParseStateChange 解析 StateChange 事件日志
func ParseStateChange(log *types.Log) (*StateChange, error) {
	var event StateChange
	if err := pledgeAbi.UnpackIntoInterface(&event, "StateChange", log.Data); err != nil {
		return nil, err
	}
	event.Pid = new(big.Int).SetBytes(log.Topics[1].Bytes())
	event.BeforeState = new(big.Int).SetBytes(log.Topics[2].Bytes())
	event.AfterState = new(big.Int).SetBytes(log.Topics[3].Bytes())
	logger.Log.Info("解析 StateChange事件成功",zap.String("pid",event.Pid.String()),zap.String("beforeState",event.BeforeState.String()),zap.String("afterState",event.AfterState.String()))
	return &event, nil


}