# EasySwap NFT Market Backend
## 项目介绍
基于 Golang 开发的 **Web3 NFT 交易市场后端服务**，
集成链上事件监听、合约ABI解析、数据持久化、区块分叉处理、定时扫链、RESTful API 能力。

采用企业标准三层架构：
`Controller(api) → Service(业务) → DAO(数据访问)`
搭配 **Viper + YAML** 配置化管理，结构清晰、解耦彻底、可直接用于简历/毕设/生产部署。

## 技术栈
- 核心语言：Go 1.21
- Web 框架：Gin
- 数据库：MySQL + GORM
- 链上交互：go-ethereum
- 配置管理：Viper(YAML)
- 架构模式：Controller / Service / DAO
- 核心能力：链上日志监听、ABI解码、定时扫链、分叉回滚、断点续扫

## 项目目录
easy-swap/
├── config.yaml # 全局 YAML 配置文件
├── go.mod
├── main.go # 项目入口
├── README.md
├── config/ # 配置初始化 + 链配置
├── dal/ # 数据库初始化
├── model/ # 数据库实体模型
├── router/ # 路由统一注册
├── api/ # Controller 控制层
├── service/ # Service 业务逻辑层
├── dao/ # DAO 数据访问层
├── internal/
│ ├── parser/ # 合约 ABI 事件解析
│ └── scan/ # 后台定时扫链服务
└── utils/ # 全局工具方法


## 核心功能
1. 链上事件监听
- 定时轮询 BSC/ETH 区块日志
- 监听 ERC721 Transfer、订单创建/取消/成交事件

2. 链上数据解析
- 通过合约 ABI 解码 topics / data
- 自动结构化入库，保证数据标准化

3. 高可用数据保障
- 区块高度持久化，服务重启**断点续扫**
- TxHash 唯一索引，天然**幂等防重**
- 自动检测**链分叉**，自动回滚分叉脏数据

4. 运维能力
- RPC 节点断线自动重连
- 手动区块回溯接口，一键重扫指定区间区块
- 全量配置 YAML 化，无需改代码切换环境

5. 业务接口
- NFT 转账记录查询
- NFT 市场在售挂单列表
- 后台运维调试接口

## 快速部署
### 1. 环境依赖
- Go 1.21+
- MySQL 5.7 / 8.0
- 可用区块链 RPC 节点(BSC/ETH)

### 2. 创建数据库
```sql
CREATE DATABASE easyswap DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
### 3. 修改配置
  -  编辑 config.yaml
  -  修改 mysql.dsn 账号密码
  -  替换 chain.rpc 为有效节点
  -  填入业务合约地址 (nft /market)
### 4. 安装依赖 & 启动
  -  bash
  -  运行
  -  go mod tidy
  -  go run main.go
  -  启动成功：
  -  HTTP 服务默认监听 :8080
  -  后台自动启动定时扫链协程
  -  自动迁移数据表结构



## 接口文档
  -   1. 获取 NFT 转账记录
  -  http GET /api/nft/transfers
  -  2. 获取在售挂单列表
  -  http GET /api/order/list
  -  3. 手动区块回溯重扫
  -  http GET /api/debug/rescan?step=100
  -  step：需要回溯多少个区块

## 架构设计说明
    分层职责
    Controller(api)
    接收 HTTP 请求
    参数校验、统一响应封装
    只调用 Service，无业务、无 DB 操作
    Service
    核心业务逻辑处理
    数据组装、业务规则判断
    只调用 DAO，不直接操作数据库
    DAO
    纯数据库 CRUD 操作
    无业务逻辑，通用数据操作复用
    Model
    数据表结构定义
    索引、字段备注、约束定义
    config+viper
    全量配置 YAML 化
    支持快速切换开发 / 生产环境
