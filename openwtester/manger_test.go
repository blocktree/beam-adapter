package openwtester

import (
	"fmt"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"testing"
)

func TestCreateLocalAddress(t *testing.T) {
	addrs, err := clientNode.CreateLocalWalletAddress(10, 2)
	if err != nil {
		t.Errorf("CreateRemoteWalletAddress failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addrs {
		log.Infof("%d: %s", i, a)
	}
}

func TestCreateRemoteAddress(t *testing.T) {
	addrs, err := clientNode.CreateRemoteWalletAddress(10, 10)
	if err != nil {
		t.Errorf("CreateRemoteWalletAddress failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addrs {
		log.Infof("%d: %s", i, a)
	}
}

func TestGetWalletBalance(t *testing.T) {
	balanceLocal, err := clientNode.GetLocalWalletBalance()
	if err != nil {
		t.Errorf("GetLocalWalletBalance failed unexpected error: %v\n", err)
		return
	}

	balanceRemote, err := clientNode.GetRemoteWalletBalance()
	if err != nil {
		t.Errorf("GetRemoteWalletBalance failed unexpected error: %v\n", err)
		return
	}

	log.Infof("local balance: %+v", balanceLocal)
	log.Infof("remote balance: %+v", balanceRemote)

}

func TestGetLocalWalletAddress(t *testing.T) {
	addrs, err := clientNode.GetLocalWalletAddress()
	if err != nil {
		t.Errorf("CreateRemoteWalletAddress failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addrs {
		log.Infof("%d: %s", i, a)
	}
}

func TestGetRemoteWalletAddress(t *testing.T) {
	addrs, err := clientNode.GetRemoteWalletAddress()
	if err != nil {
		t.Errorf("CreateRemoteWalletAddress failed unexpected error: %v\n", err)
		return
	}

	for _, a := range addrs {
		fmt.Printf("%s\n", a)
	}
}

func TestGetRemoteBlockByHeight(t *testing.T) {
	block, err := clientNode.GetRemoteBlockByHeight(313212)
	if err != nil {
		t.Errorf("CreateRemoteWalletAddress failed unexpected error: %v\n", err)
		return
	}

	log.Infof("%+v", block)
}

func TestSendTransaction(t *testing.T) {

	rawTx := &openwallet.RawTransaction{
		Coin:openwallet.Coin{
			Symbol:"BEAM",
		},
		To: map[string]string{
			"3c4848707c2ce5e83c48915ecbccd8e96aa660bae4992b7f7ef7373c266d83aee46": "0.0000005",
		},
		FeeRate: "",
	}

	txdecoder := clientNode.TxDecoder
	//err := txdecoder.CreateRawTransaction(nil, rawTx)
	//if err != nil {
	//	t.Errorf("CreateRawTransaction failed unexpected error: %v\n", err)
	//	return
	//}
	//
	//err = txdecoder.SignRawTransaction(nil, rawTx)
	//if err != nil {
	//	t.Errorf("SignRawTransaction failed unexpected error: %v\n", err)
	//	return
	//}
	//
	//err = txdecoder.VerifyRawTransaction(nil, rawTx)
	//if err != nil {
	//	t.Errorf("VerifyRawTransaction failed unexpected error: %v\n", err)
	//	return
	//}

	tx, err := txdecoder.SubmitRawTransaction(nil, rawTx)
	if err != nil {
		t.Errorf("SubmitRawTransaction failed unexpected error: %v\n", err)
		return
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)
}

func TestGetBalance(t *testing.T) {
	scanner := clientNode.Blockscanner
	balance, err := scanner.GetBalanceByAddress()
	if err != nil {
		t.Errorf("GetBalanceByAddress failed unexpected error: %v\n", err)
		return
	}
	log.Infof("balance: %+v", balance[0])
}

func TestStartSummaryWallet(t *testing.T) {

	serverNode.StartSummaryWallet()

}