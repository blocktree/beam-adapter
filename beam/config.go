package beam

import (
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/common/file"
	"path/filepath"
	"strings"
)

const (
	//币种
	Symbol    = "BEAM"
	CurveType = owcrypt.ECC_CURVE_SECP256K1

	//默认配置内容
	defaultConfig = `

# Beam Wallet RPC API
walletapi = "http://47.91.224.127:20021/api/wallet"

# Beam explore API
explorerapi = "http://47.91.224.127:20022"

# beam-adapter Remote Server
remoteserver = "127.0.0.1"

# True: Run for server, False: Run for client
enableserver = false

# Fix Transaction Fess
fixfees = "0.1"

# Node Connect Type
connecttype = "websocket"

# Enable key agreement on local node communicate with client server
enablekeyagreement = false

# Enable https or wss
enablessl = false

# Network request timeout, unit: second
requesttimeout = 120

# Generate Node
cert = ""

# trust node id
trustnodeid = ""

# summary address 汇总地址
summaryaddress = ""

# summary threshold 汇总阈值
summarythreshold = ""

# Wallet Summary Period,  汇总周期
summaryperiod = "30s"

# Log file path
logdir = "./logs/"

# log debug info
logdebug = false

`
)

type WalletConfig struct {

	//币种
	Symbol string
	//配置文件路径
	configFilePath string
	//配置文件名
	configFileName string
	//区块链数据文件
	BlockchainFile string
	//本地数据库文件路径
	dbPath string
	//默认配置内容
	DefaultConfig string
	//曲线类型
	CurveType uint32
	//固定手续费
	fixfees string
	// 远程服务
	remoteserver string
	//是否开启协商密码通信
	enablekeyagreement bool
	//是否支持ssl：https，wss等
	enablessl bool
	//网络请求超时，单位：秒
	requesttimeout int
	//钱包API
	walletapi string
	//浏览器API
	explorerapi string
	//连接方式
	connecttype string
	//信任节点
	trustnodeid string
	//是否作为服务端
	enableserver bool
	//是否输出LogDebugg日志
	logdebug bool
	//通信证书私钥
	cert string
	//汇总地址
	summaryaddress string
	//汇总阈值
	summarythreshold string
	//汇总时间周期
	summaryperiod string
	//日志路径
	logdir string
}

func NewConfig(symbol string) *WalletConfig {

	c := WalletConfig{}

	//币种
	c.Symbol = symbol
	c.CurveType = CurveType

	//区块链数据
	//blockchainDir = filepath.Join("data", strings.ToLower(Symbol), "blockchain")
	//配置文件路径
	c.configFilePath = filepath.Join("conf")
	//配置文件名
	c.configFileName = c.Symbol + ".ini"
	//区块链数据文件
	c.BlockchainFile = "blockchain.db"
	//本地数据库文件路径
	c.dbPath = filepath.Join("data", strings.ToLower(c.Symbol), "db")

	//创建目录
	file.MkdirAll(c.dbPath)

	return &c
}
