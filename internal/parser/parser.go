package parser

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	nftAbi    abi.ABI
	marketAbi abi.ABI
)

// Init 初始化abi解析器
func Init() error {
	// NFT Transfer事件ABI
	nftABI := `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"}]`
	parsedNFT, err := abi.JSON(strings.NewReader(nftABI))
	if err != nil {
		return err
	}
	nftAbi = parsedNFT

	// 市场订单事件ABI
	marketABI := `[
	{"anonymous":false,"inputs":[{"indexed":true,"name":"seller","type":"address"},{"indexed":false,"name":"orderId","type":"uint256"},{"indexed":false,"name":"tokenId","type":"uint256"},{"indexed":false,"name":"price","type":"uint256"}],"name":"OrderCreated","type":"event"},
	{"anonymous":false,"inputs":[{"indexed":false,"name":"orderId","type":"uint256"}],"name":"OrderCancelled","type":"event"},
	{"anonymous":false,"inputs":[{"indexed":false,"name":"orderId","type":"uint256"},{"indexed":true,"name":"buyer","type":"address"}],"name":"OrderFilled","type":"event"}
]`
	parsedMarket, err := abi.JSON(strings.NewReader(marketABI))
	if err != nil {
		return err
	}
	marketAbi = parsedMarket
	return nil
}

// TransferEvent NFT转账事件
type TransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId string
}

// OrderCreatedEvent 订单创建事件
type OrderCreatedEvent struct {
	Seller  common.Address
	OrderId string
	TokenId string
	Price   string
}

// OrderCancelledEvent 订单取消事件
type OrderCancelledEvent struct {
	OrderId string
}

// OrderFilledEvent 订单成交事件
type OrderFilledEvent struct {
	OrderId string
	Buyer   common.Address
}

// ParseTransfer 解析NFT转账日志
func ParseTransfer(log *types.Log) (*TransferEvent, error) {
	var event TransferEvent
	if err := nftAbi.UnpackIntoInterface(&event, "Transfer", log.Data); err != nil {
		return nil, err
	}
	event.From = common.BytesToAddress(log.Topics[1].Bytes())
	event.To = common.BytesToAddress(log.Topics[2].Bytes())
	return &event, nil
}

// ParseOrderCreated 解析创建订单日志
func ParseOrderCreated(log *types.Log) (*OrderCreatedEvent, error) {
	var event OrderCreatedEvent
	if err := marketAbi.UnpackIntoInterface(&event, "OrderCreated", log.Data); err != nil {
		return nil, err
	}
	event.Seller = common.BytesToAddress(log.Topics[1].Bytes())
	return &event, nil
}

// ParseOrderCancelled 解析取消订单日志
func ParseOrderCancelled(log *types.Log) (*OrderCancelledEvent, error) {
	var event OrderCancelledEvent
	if err := marketAbi.UnpackIntoInterface(&event, "OrderCancelled", log.Data); err != nil {
		return nil, err
	}
	return &event, nil
}

// ParseOrderFilled 解析订单成交日志
func ParseOrderFilled(log *types.Log) (*OrderFilledEvent, error) {
	var event OrderFilledEvent
	if err := marketAbi.UnpackIntoInterface(&event, "OrderFilled", log.Data); err != nil {
		return nil, err
	}
	event.Buyer = common.BytesToAddress(log.Topics[1].Bytes())
	return &event, nil
}