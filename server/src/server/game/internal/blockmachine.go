package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
	"server/mysql"
	"github.com/name5566/leaf/log"
	"github.com/opesun/goquery"
	"strings"
	"fmt"
)

//仓库里的矿机数据连，这里没有进数据库，待处理

// 已购买的矿机
type BlockMachine struct {
	Index     int
	Timestamp int64
	PlayerID  uint
	DataID    int
	Price     float32 //单价
	Count     int     //数量
	Hash      string
	PrevHash  string
}

var DepotChain []BlockMachine
var mutex = &sync.Mutex{}
var PeerDayBtc float32 //每天收益(1TH/s)
func DepotChainInit() {

	//先从数据库读取，没有的话再初始化
	db := mysql.MysqlDB()
	rows, err := db.Query("select `index`, `timestamp`, `player_id`,`data_id`,`price`,`count`,`hash`,`prev_hash` from block_machines")
	if nil == err {
		for rows.Next() {
			u := BlockMachine{}
			err = rows.Scan(&u.Index, &u.Timestamp, &u.PlayerID, &u.DataID, &u.Price, &u.Count, &u.Hash, &u.PrevHash)
			if nil == err {
				DepotChain = append(DepotChain, u)
			}
		}
	}
	mysql.CheckErr(err)
	if len(DepotChain) > 0 {
		return
	}

	//没有才初始化
	t := time.Now().Unix()
	genesisBlock := BlockMachine{}
	genesisBlock = BlockMachine{0, t, 0, 0, 0, 0, calculateHash(genesisBlock), ""}
	//spew.Dump(genesisBlock)

	mutex.Lock()
	DepotChain = append(DepotChain, genesisBlock)
	mutex.Unlock()



}
func GetBlockMachinechain(playerID uint) {
	var playerDepotChain []BlockMachine
	for _, r := range DepotChain {
		playerDepotChain = append(playerDepotChain, r)
	}
}

//得到某个矿机的库存量
func GetRepertory(productID int, maxCount int) int {
	value := 0
	for _, r := range DepotChain {
		if (r.DataID == productID) {
			value ++
		}
	}
	return (maxCount - value)
}

func WriteBlockMachine(playerID uint, dataID int, price float32, count int) BlockMachine {

	mutex.Lock()
	newBlock := generateBlock(DepotChain[len(DepotChain)-1], playerID, dataID, price, count)
	mutex.Unlock()

	if isBlockValid(newBlock, DepotChain[len(DepotChain)-1]) {
		DepotChain = append(DepotChain, newBlock)
		//spew.Dump(DepotChain)

		db := mysql.MysqlDB()

		stmt, err := db.Prepare("insert block_machines values(?,?,?,?,?,?,?,?)")
		if nil == err {
			_, err = stmt.Exec(newBlock.Index, newBlock.Timestamp, newBlock.PlayerID, newBlock.DataID, newBlock.Price, newBlock.Count, newBlock.Hash, newBlock.PrevHash)
		}
		if nil != err {
			log.Error("create WriteBlockMachine error:", err)
		}
	}

	return newBlock

}

// make sure block is valid by checking index, and comparing the hash of the previous block
func isBlockValid(newBlock, oldBlock BlockMachine) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// SHA256 hasing
func calculateHash(block BlockMachine) string {
	record := strconv.Itoa(block.Index) + string(block.Timestamp) + string(block.PlayerID) + string(block.DataID) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func generateBlock(oldBlock BlockMachine, playerID uint, dataID int, price float32, count int) BlockMachine {

	var newBlock BlockMachine
	t := time.Now().Unix()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t
	newBlock.PlayerID = playerID
	newBlock.DataID = dataID
	newBlock.Price = price
	newBlock.Count = count
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

//当前PPS+理论每天收益(1TH/s)
func UpdatePeerDayBtc() {

	var url = "https://www.antpool.com/earnComparison.htm"
	p, err := goquery.ParseUrl(url)
	if err != nil {
		panic(err)
	} else {
		t := p.Find("span").Text()
		startpos := strings.Index(t, "BTC:0.")
		endpos := strings.Index(t, "BCH:0.")

		btcStr := t[startpos+4 : endpos]
		vtc, _ := strconv.ParseFloat(btcStr, 32)
		PeerDayBtc = float32(vtc)
		log.Release("PeerDayBtc:%v", strconv.FormatFloat(vtc, 'f', -1, 32))
		fmt.Println("PeerDayBtc:%v",strconv.FormatFloat(vtc, 'f', -1, 32))
	}
}

//每分钟更新
func OnTickMinute() {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				//fmt.Printf("ticked at %v\n", time.Now())
				UpdatePeerDayBtc()
			}
		}

	}()
}

//每小时更新
func OnTickHour() {
	ticker := time.NewTicker(time.Hour * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				db := mysql.MysqlDB()
				db.Ping()

				UpdatePeerDayBtc()
			}
		}
	}()
}
