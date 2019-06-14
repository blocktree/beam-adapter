# beam-adapter

beam暂时无法适配openwallet钱包体系，无法实现离线生产地址，离线交易签名。
因为beam的Mimblewimble协议需要钱包一直保持在线，才能正常地接收和发送资产。
为了实现企业级别钱包解决方案，我们划分两个应用场景。
- 用户托管钱包场景。就是企业托管了注册用户的钱包，为每一个用户分配接收地址，这个钱包不对外开放。在一台装有beam钱包的服务器上，运行beam-wallet钱包服务，监听用户的充值，并定时进行汇总。
- 财务系统提币钱包场景。企业不会直接从用户托管钱包中做转账，而是在一个独立的热钱包中驻留日常业务的需要的资产。财务系统集成beam-adapter，可发送指令向beam-wallet钱包服务创建新地址，可查询用户托管钱包的新充值记录。

## 官方资料

### 官网

https://www.beam.mw/

### 接口文档

#### Wallet API

https://github.com/BeamMW/beam/wiki/Beam-wallet-protocol-API

#### Explorer API

https://github.com/BeamMW/beam/wiki/Beam-Node-Explorer-API

### 浏览器

https://explorer.beam.mw/

### 测试币领取

https://bitmate.ch/

## beam-wallet使用说明

beam-wallet是一个运行在beam用户托管钱包服务器的后台程序，为财务系统的热钱包提供创建用户地址，获取交易记录，获取钱包余额，执行定时汇总等功能。

### 编译程序

#### 安装xgo（支持跨平台编译C代码）

[官方github（目前还不支持go module）](https://github.com/karalabe/xgo)
[支持go module的xgo fork](https://github.com/gythialy/xgo)

xgo的使用依赖docker。并且把要跨平台编译的项目文件加入到File sharing。

```shell

# 官方的目前还不支持go module编译，所以我们用了别人改造后能够给支持的fork版
$ go get -u github.com/gythialy/xgo
...
$ xgo -h
...

# 本地系统编译
$ make clean build

# 跨平台编译wmd，更多可选平台可修改Makefile的$TARGETS变量
$ make clean beam-wallet


```

### 服务端配置文件

在用户托管钱包的服务器配置，运行beam-wallet后台服务，开放接口给财务系统的热钱包使用，样例如下：

```ini

# Beam Wallet RPC API, beam钱包API
walletapi = "http://192.168.1.123:12345/api/wallet"

# Beam explore API, beam钱包浏览器API
explorerapi = "http://192.168.1.123:12346"

# beam-adapter Remote Server, beam-wallet服务的固定IP或域名
remoteserver = "127.0.0.1:20888"

# True: Run for server, False: Run for client, 作为服务端启动
enableserver = true

# Fix Transaction Fess, 最低手续费
fixfees = "0.00000001"

# Node Connect Type, 连接方式：ws: websocket
connecttype = "ws"

# Enable key agreement on local node communicate with client server, 开启协商密码
enablekeyagreement = true

# log debug info, 是否打印debug日志
logdebug = false

# Log file path, 日志目录
logdir = "./logs/"

# trust node id, 服务端让授信的客户端连接
trustnodeid = "11111"

# summary address 汇总地址
summaryaddress = "111111"

# summary threshold 汇总阈值
summarythreshold = "0.001"

# Wallet Summary Period,  汇总周期
summaryperiod = "30s"

# Transaction sending timeout, 如果接受方钱包不在线，交易会一直处于发送中状态，需要设置一个超时时间，超时取消发送中的交易
# Such as "30s", "1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
txsendingtimeout = "5m"

```

在用户托管钱包的服务器运行beam-walle

```shell

# 加载配置server.ini，运行walletserver后台服务
$ ./beam-wallet -c=server.ini walletserver

```

### 客户端配置文件

在财务系统的钱包服务器配置，财务系统集成beam-adapter，通过AssetsAdapter接口加载如下配置：

```ini

# Beam Wallet RPC API, beam钱包API
walletapi = "http://192.168.1.123:12345/api/wallet"

# Beam explore API, beam钱包浏览器API
explorerapi = "http://192.168.1.123:12346"

# beam-adapter Remote Server, beam-wallet服务的固定IP或域名
remoteserver = "127.0.0.1:20888"

# True: Run for server, False: Run for client, 作为客户端启动
enableserver = false

# Fix Transaction Fess, 最低手续费
fixfees = "0.00000001"

# Node Connect Type, 连接方式：ws: websocket
connecttype = "ws"

# Enable key agreement on local node communicate with client server, 开启协商密码
enablekeyagreement = true

# Enable https or wss, 如果服务器有SSL证书，可开启SSL，
enablessl = false

# Network request timeout, unit: second, 请求连接超时限制
requesttimeout = 120

# log debug info, 是否打印debug日志
logdebug = false

# Log file path, 日志目录
logdir = "./logs/"

# Generate Node, 客户端证书私钥
cert = "1111"

# Transaction sending timeout, 如果接受方钱包不在线，交易会一直处于发送中状态，需要设置一个超时时间，超时取消发送中的交易
# Such as "30s", "1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
txsendingtimeout = "5m"

```

系统集成beam-adapter/beam包功能

```go

    //创建beam钱包管理对象
	clientNode := beam.NewWalletManager()
	
	//加载client.ini配置文件
	c, err := config.NewConfig("ini", "client.ini")
	if err != nil {
		return nil
	}
	clientNode.LoadAssetsConfig(c)
	
	//向远程服务，创建用户托管钱包的地址
	addrs, err := clientNode.CreateRemoteWalletAddress(100, 10)
	if err != nil {
        return
	}
	
	//发起转账交易
    rawTx := &openwallet.RawTransaction{
        To: map[string]string{
            "3b769e29f6e2fc59fb7d1cd88fa03bd0777318b83d0e5111941992ad5efbe670d31": "0.0000001",
        },
        FeeRate: "",
    }

    txdecoder := clientNode.TxDecoder
    tx, err := txdecoder.SubmitRawTransaction(nil, rawTx)
    
    //启动区块链扫描器
    scanner := clientNode.GetBlockScanner()
	scanner.Run()
	
```