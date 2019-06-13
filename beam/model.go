package beam

import (
	"fmt"
	"github.com/blocktree/openwallet/common"
	"github.com/blocktree/openwallet/crypto"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/tidwall/gjson"
)

type Block struct {
	Chainwork     string
	Hash          string
	Found         bool
	PrevBlockHash string
	Time          int64
	Height        uint64
	inputs        []interface{}
	kernels       []interface{}
	outputs       []interface{}

	/*
		{
		  "chainwork": "0x384fdc718a20",
		  "difficulty": 157.9972152709961,
		  "found": true,
		  "hash": "c2a7315b63b1de6106a185c1c79219001ef5e3a07c217db227b079bbb9dd9b64",
		  "height": 20516,
		  "inputs": [],
		  "kernels": [
		    {
		      "excess": "0x60413b5a09858312403190721938463ca22d7d87981a024873ddfa204a399eec",
		      "fee": 0,
		      "id": "d72684dba6255b2fe8631be9df764ee3c984cb0c9f386a8cf71f566acebd197d",
		      "maxHeight": 18446744073709552000,
		      "minHeight": 20516
		    }
		  ],
		  "outputs": [
		    {
		      "coinbase": true,
		      "commitment": "0x75329071d041e7828a57cbf2f63fb8db21543f35c1c2291d5c26c20d9b11465a",
		      "incubation": 0,
		      "maturity": 20756
		    }
		  ],
		  "prev": "4b9e35b467b416e0d307dd94bd2fdce6e720b6b3a029dca822ccab3ac57c6d22",
		  "subsidy": 8000000000,
		  "timestamp": 1550157362
		}

	*/
}

func NewBlock(result *gjson.Result) *Block {
	obj := Block{}
	//解析json
	obj.Hash = result.Get("hash").String()
	obj.Chainwork = result.Get("chainwork").String()
	obj.PrevBlockHash = result.Get("prev").String()
	obj.Height = result.Get("height").Uint()
	obj.Time = result.Get("timestamp").Int()
	obj.Found = result.Get("found").Bool()

	return &obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader(symbol string) *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	//obj.Confirmations = b.Confirmations
	//obj.Merkleroot = b.TransactionMerkleRoot
	obj.Previousblockhash = b.PrevBlockHash
	obj.Height = b.Height
	obj.Time = uint64(obj.Time)
	obj.Symbol = symbol

	return &obj
}

type Transaction struct {
	Comment       string
	CreateTime    int64
	Fee           uint64
	TxID          string
	Value         uint64
	Kernel        string
	Receiver      string
	Sender        string
	Income        bool
	Status        int64
	StatusString  string
	Confirmations uint64
	BlockHeight   uint64
	BlockHash     string

	/*
			{
		        "comment": "",
		        "confirmations": 5,
		        "create_time": 1559873867,
		        "fee": 1,
		        "height": 221412,
		        "income": true,
		        "kernel": "cf3634952569171015fe08b949ed617692a30947747fb576e826d6f48a1b8035",
		        "receiver": "21aff5eb4da2591321ac12bb280ac69ea39a33472166c600ec122cf3381b6c9e772",
		        "sender": "22d090004ab6de7e62d0d3829e0164d05cc065404ebc9874d181dc070d54237bbd8",
		        "status": 3,
		        "status_string": "received",
		        "txId": "72f8f349f9244b11b0e6471250ca68a1",
		        "value": 10000
		    }
	*/
}

func NewTransaction(result *gjson.Result) *Transaction {
	obj := Transaction{}
	obj.Comment = result.Get("comment").String()
	obj.CreateTime = result.Get("create_time").Int()
	obj.Fee = result.Get("fee").Uint()
	obj.Income = result.Get("income").Bool()
	obj.Kernel = result.Get("kernel").String()
	obj.Receiver = result.Get("receiver").String()
	obj.Sender = result.Get("sender").String()
	obj.Status = result.Get("status").Int()
	obj.StatusString = result.Get("status_string").String()
	obj.TxID = result.Get("txId").String()
	obj.Value = result.Get("value").Uint()
	obj.Confirmations = result.Get("confirmations").Uint()
	obj.BlockHeight = result.Get("height").Uint()

	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}

type TrustNodeInfo struct {
	NodeID      string `json:"nodeID"` //@required 节点ID
	NodeName    string `json:"nodeName"`
	ConnectType string `json:"connectType"`
}

type BlockchainInfo struct {
	Chainwork  string
	Hash       string
	Height     uint64
	LowHorizon uint64
	Timestamp  int64

	/*
		{
			"chainwork": "0x38594101d0a0",
			"hash": "7353b5e4ad29a2ffa5f7952749d1eb04acedd82215b1f4f01d75107165f4622b",
			"height": 20531,
			"low_horizon": 19090,
			"timestamp": 1550158283
		}
	*/
}

func NewBlockchainInfo(result *gjson.Result) *BlockchainInfo {
	obj := BlockchainInfo{}
	obj.Chainwork = result.Get("H.chainwork").String()
	obj.Hash = result.Get("hash").String()
	obj.Height = result.Get("height").Uint()
	obj.LowHorizon = result.Get("low_horizon").Uint()
	obj.Timestamp = result.Get("timestamp").Int()
	return &obj
}

type WalletStatus struct {
	CurrentHeight    uint64
	CurrentStateHash string
	PrevStateHash    string
	Available        uint64
	Receiving        uint64
	Sending          uint64
	Maturing         uint64
	Locked           uint64

	/*
		{
		    "current_height" : 1055,
		    "current_state_hash" : "f287176bdd517e9c277778e4c012bf6a3e687dd614fc552a1ed22a3fee7d94f2",
		    "prev_state_hash" : "bd39333a66a8b7cb3804b5978d42312c841dbfa03a1c31fc2f0627eeed6e43f2",
		    "available": 100500,
		    "receiving": 123,
		    "sending": 0,
		    "maturing": 50,
		    "locked": 30,
		    "difficulty": 2.93914,
		}
	*/
}

func NewWalletStatus(result *gjson.Result) *WalletStatus {
	obj := WalletStatus{}
	obj.CurrentHeight = result.Get("current_height").Uint()
	obj.CurrentStateHash = result.Get("current_state_hash").String()
	obj.PrevStateHash = result.Get("prev_state_hash").String()
	obj.Available = result.Get("available").Uint()
	obj.Receiving = result.Get("receiving").Uint()
	obj.Sending = result.Get("sending").Uint()
	obj.Maturing = result.Get("maturing").Uint()
	obj.Locked = result.Get("locked").Uint()
	return &obj
}

type AddressCreateResult struct {
	Success bool
	Err     error
	Address string
}