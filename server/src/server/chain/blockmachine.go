package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
	"github.com/davecgh/go-spew/spew"
)

//仓库里的矿机数据连，这里没有进数据库，待处理

// 已购买的矿机
type BlockMachine struct {
	Index     int
	Timestamp string
	PlayerID  uint
	DataID    int
	Hash      string
	PrevHash  string
}

var DepotChain []BlockMachine
var mutex = &sync.Mutex{}

func Init() {

	t := time.Now()
	genesisBlock := BlockMachine{}
	genesisBlock = BlockMachine{0, t.String(), 0, 0, calculateHash(genesisBlock), ""}
	//spew.Dump(genesisBlock)

	mutex.Lock()
	DepotChain = append(DepotChain, genesisBlock)
	mutex.Unlock()

}
func GetBlockMachinechain(PlayerID uint)  {
	var playerDepotChain []BlockMachine
	for _, r := range DepotChain {
		playerDepotChain = append(playerDepotChain, r)
	}
}

//得到某个矿机的库存量
func GetRepertory(data int, maxCount int) int {
	value := 0
	for _, r := range DepotChain {
		if (r.DataID == data){
			value ++
		}
	}
	return (maxCount-value)
}

func WriteBlockMachine(PlayerID uint, DataID int) BlockMachine {

	mutex.Lock()
	newBlock := generateBlock(DepotChain[len(DepotChain)-1], PlayerID, DataID)
	mutex.Unlock()

	if isBlockValid(newBlock, DepotChain[len(DepotChain)-1]) {
		DepotChain = append(DepotChain, newBlock)
		spew.Dump(DepotChain)
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
	record := strconv.Itoa(block.Index) + block.Timestamp + string(block.PlayerID) + string(block.DataID) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func generateBlock(oldBlock BlockMachine, PlayerID uint, DataID int) BlockMachine {

	var newBlock BlockMachine

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PlayerID = PlayerID
	newBlock.DataID = DataID
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
