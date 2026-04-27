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
	nftABI := `[    {
		"anonymous": false,
		"inputs": [
		  {
			"indexed": true,
			"name": "from",
			"type": "address"
		  },
		  {
			"indexed": true,
			"name": "to",
			"type": "address"
		  },
		  {
			"indexed": true,
			"name": "tokenId",
			"type": "uint256"
		  }
		],
		"name": "Transfer",
		"type": "event"
	  }]`
	parsedNFT, err := abi.JSON(strings.NewReader(nftABI))
	if err != nil {
		return err
	}
	nftAbi = parsedNFT

	// 市场订单事件ABI
	marketABI := `[{
		"anonymous": false,
		"inputs": [
		  {
			"indexed": false,
			"name": "orderKey",
			"type": "bytes32"
		  },
		  {
			"indexed": true,
			"name": "side",
			"type": "uint8"
		  },
		  {
			"indexed": true,
			"name": "saleKind",
			"type": "uint8"
		  },
		  {
			"indexed": true,
			"name": "maker",
			"type": "address"
		  },
		  {
			"components": [
			  {
				"name": "tokenId",
				"type": "uint256"
			  },
			  {
				"name": "collection",
				"type": "address"
			  },
			  {
				"name": "amount",
				"type": "uint96"
			  }
			],
			"indexed": false,
			"name": "nft",
			"type": "tuple"
		  },
		  {
			"indexed": false,
			"name": "price",
			"type": "uint128"
		  },
		  {
			"indexed": false,
			"name": "expiry",
			"type": "uint64"
		  },
		  {
			"indexed": false,
			"name": "salt",
			"type": "uint64"
		  }
		],
		"name": "LogMake",
		"type": "event"
	  }, {
		"anonymous": false,
		"inputs": [
		  {
			"indexed": true,
			"name": "orderKey",
			"type": "bytes32"
		  },
		  {
			"indexed": true,
			"name": "maker",
			"type": "address"
		  }
		],
		"name": "LogCancel",
		"type": "event"
	  },  {
		"anonymous": false,
		"inputs": [
		  {
			"indexed": true,
			"name": "makeOrderKey",
			"type": "bytes32"
		  },
		  {
			"indexed": true,
			"name": "takeOrderKey",
			"type": "bytes32"
		  },
		  {
			"components": [
			  {
				"name": "side",
				"type": "uint8"
			  },
			  {
				"name": "saleKind",
				"type": "uint8"
			  },
			  {
				"name": "maker",
				"type": "address"
			  },
			  {
				"components": [
				  {
					"name": "tokenId",
					"type": "uint256"
				  },
				  {
					"name": "collection",
					"type": "address"
				  },
				  {
					"name": "amount",
					"type": "uint96"
				  }
				],
				"name": "nft",
				"type": "tuple"
			  },
			  {
				"name": "price",
				"type": "uint128"
			  },
			  {
				"name": "expiry",
				"type": "uint64"
			  },
			  {
				"name": "salt",
				"type": "uint64"
			  }
			],
			"indexed": false,
			"name": "makeOrder",
			"type": "tuple"
		  },
		  {
			"components": [
			  {
				"name": "side",
				"type": "uint8"
			  },
			  {
				"name": "saleKind",
				"type": "uint8"
			  },
			  {
				"name": "maker",
				"type": "address"
			  },
			  {
				"components": [
				  {
					"name": "tokenId",
					"type": "uint256"
				  },
				  {
					"name": "collection",
					"type": "address"
				  },
				  {
					"name": "amount",
					"type": "uint96"
				  }
				],
				"name": "nft",
				"type": "tuple"
			  },
			  {
				"name": "price",
				"type": "uint128"
			  },
			  {
				"name": "expiry",
				"type": "uint64"
			  },
			  {
				"name": "salt",
				"type": "uint64"
			  }
			],
			"indexed": false,
			"name": "takeOrder",
			"type": "tuple"
		  },
		  {
			"indexed": false,
			"name": "fillPrice",
			"type": "uint128"
		  }
		],
		"name": "LogMatch",
		"type": "event"
	  }]`
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
