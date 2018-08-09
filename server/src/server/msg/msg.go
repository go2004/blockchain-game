package msg
import (
	"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

func init() {	// 这里我们注册 protobuf 消息)
    Processor.SetByteOrder(true)
    Processor.Register(&U2S_Login{})
    Processor.Register(&S2U_LoginResult{})
    Processor.Register(&U2S_RegistAccount{})
    Processor.Register(&G2U_RegistAccountResult{})
    Processor.Register(&S2U_PlayerBase{})
    Processor.Register(&S2U_Notice{})
    Processor.Register(&U2S_Chat{})
    Processor.Register(&S2U_Chat{})
    Processor.Register(&U2S_ShopList{})
    Processor.Register(&ShopInfo{})
    Processor.Register(&S2U_ShopList{})
    Processor.Register(&U2S_Buy{})
    Processor.Register(&S2U_BuyResult{})
    Processor.Register(&U2S_MineMachineList{})
    Processor.Register(&MineMachine{})
    Processor.Register(&S2U_MineMachineList{})
    Processor.Register(&U2S_MineList{})
    Processor.Register(&MineInfo{})
    Processor.Register(&S2U_MineList{})
    Processor.Register(&U2S_Mining{})
    Processor.Register(&S2U_MiningResult{})
    Processor.Register(&U2S_GatherMine{})
    Processor.Register(&S2U_GatherMine{})
    Processor.Register(&U2S_StopMining{})
    Processor.Register(&S2U_StopMiningResult{})
    Processor.Register(&U2S_ModifyPlayer{})
    Processor.Register(&S2U_ModifyPlayer{})
    Processor.Register(&U2S_Rank{})
    Processor.Register(&RankInfo{})
    Processor.Register(&S2U_AppCoinRank{})
    Processor.Register(&S2U_BtcRank{})

}