package gate

import (
	"server/game"
	"server/msg"
	"server/login"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.U2S_Chat{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_RegistAccount{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_ShopList{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_Buy{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_MineMachineList{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_Mining{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_GatherMine{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_MineList{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_StopMining{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_ModifyPlayer{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.U2S_Rank{}, game.ChanRPC)
}
