package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/beam-adapter/beam"
	"path/filepath"
	"time"
)

var (
	serverNode *beam.WalletManager
	clientNode *beam.WalletManager
)

func init() {

	serverNode = testNewWalletManager("server.ini")
	clientNode = testNewWalletManager("client.ini")
	time.Sleep(1 * time.Second)
}

func testNewWalletManager(conf string) *beam.WalletManager {
	wm := beam.NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", conf)
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	return wm
}

