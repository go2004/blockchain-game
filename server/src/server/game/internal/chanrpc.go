package internal

import (
	"github.com/name5566/leaf/gate"
	"server/msg"
	"github.com/name5566/leaf/log"
	"server/common"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("UserLogin", rpcUserLogin)
}

func rpcNewAgent(args []interface{}) {
	//fmt.Println("rpcNewAgent")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	//fmt.Println("rpcCloseAgent:1")
	a := args[0].(gate.Agent)
	//fmt.Println("rpcCloseAgent:2")
	//fmt.Println("rpcCloseAgent:2 ->",a.UserData().(uint))
	if a.UserData() != nil {
		playerID := a.UserData().(uint)
		player := playerID2Player[playerID]
		player.state = userGame
		player.onLogout()
		delete(playerID2Player,playerID)
		//fmt.Println("rpcCloseAgent:2->",playerID)
	}

	_ = a
}

func rpcUserLogin(args []interface{}) {
	agent := args[0].(gate.Agent)
	playerID := args[1].(uint)
	account := args[2].(string)
	channelID := args[3].(uint32)
	name := args[4].(string)
	style := args[5].(uint32)
	appCoins := args[6].(float32)
	computePower := args[7].(float32)
	btc := args[8].(float32)
	// network closed
	if agent.UserData() == nil {
		return
	}

	oldUser := playerID2Player[playerID]
	if oldUser != nil {
		m := &msg.S2U_LoginResult{Result: 3}
		agent.WriteMsg(m)
		oldUser.WriteMsg(m)
		agent.Close()
		oldUser.Close()
		log.Debug("acc %v login repeated", playerID)
		delete(playerID2Player,playerID)
		return
	}
	log.Debug("acc %v login", playerID)
	//fmt.Println("rpcUserLogin: ", playerID)
	// login
	newPlayer := new(Player)
	newPlayer.Agent = agent
	newPlayer.LinearContext = skeleton.NewLinearContext()
	newPlayer.state = userLogin
	//newPlayer.UserData().(*PlayerBaseInfo).PlayerID = playerID

	newPlayer.account = account
	newPlayer.channelID = channelID
	newPlayer.playerID = playerID
	newPlayer.name = name
	newPlayer.style = style
	newPlayer.appCoins = appCoins
	newPlayer.computePower = computePower
	newPlayer.logintime = common.GetTimestamp()
	newPlayer.btc = btc
	playerID2Player[playerID] = newPlayer
	newPlayer.login(playerID)
}

