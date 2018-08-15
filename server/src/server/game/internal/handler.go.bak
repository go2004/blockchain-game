package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
	"server/gamedata"

	"server/common"
	"fmt"
)

func init() {
	handler(&msg.U2S_Chat{}, handleTosChat)
	handler(&msg.U2S_ShopList{}, handleU2S_ShopList)
	handler(&msg.U2S_Buy{}, handleU2S_Buy)
	handler(&msg.U2S_MineMachineList{}, handleU2S_MineMachineList)
	handler(&msg.U2S_Mining{}, handleU2S_Mining)
	handler(&msg.U2S_GatherMine{}, handleU2S_GatherMine)
	handler(&msg.U2S_MineList{}, handleU2S_MineList)
	handler(&msg.U2S_StopMining{}, handleU2S_StopMining)
	handler(&msg.U2S_ModifyPlayer{}, handleU2S_ModifyPlayer)
	handler(&msg.U2S_Rank{}, handleU2S_Rank)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//把聊天
func handleTosChat(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.U2S_Chat)
	// 消息的发送者
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}

	//直接全服广播，暂不保存及字符验证处理 todo
	playerID := a.UserData().(uint)
	fromPlayer := playerID2Player[playerID]
	for _, p := range playerID2Player {
		p.WriteMsg(&msg.S2U_Chat{Name: fromPlayer.name, Style: fromPlayer.style, Content: m.Content,})
	}

}

//请求商城列表
func handleU2S_ShopList(args []interface{}) {
	//_ = args[0].(*msg.U2S_ShopList)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}
	onShopList(a)
}
func onShopList(a gate.Agent) {
	sndmsg := &msg.S2U_ShopList{}
	shoplist := gamedata.ShopData
	for _, r := range shoplist {
		p1 := &msg.ShopInfo{
			ProductID:  int32(r.ProductID),
			Repertory:  int32(GetRepertory(r.ProductID, r.Repertory)),
			Electric:   r.Electric,
			ManageCost: r.ManageCost,
			Output:     r.Output,
			Price:      r.Price,
		}
		sndmsg.List = append(sndmsg.List, p1)
	}
	a.WriteMsg(sndmsg)
}

//购买矿机
func handleU2S_Buy(args []interface{}) {
	m := args[0].(*msg.U2S_Buy)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}

	productID := int(m.ProductID)
	count := int(m.Count)

	//检查数据
	record, exists := gamedata.GetShopByID(productID)
	if (!exists) {
		a.WriteMsg(&msg.S2U_BuyResult{ErrorCode: 1})
		return
	}

	//数量
	freeCount := GetRepertory(productID, record.Repertory)
	if (count < 0) && (freeCount > 0) && (freeCount < count) {
		a.WriteMsg(&msg.S2U_BuyResult{ErrorCode: 2})
		return
	}

	//检查游戏币是否足够
	var total float32
	total = record.Price * float32(count)
	playerID := a.UserData().(uint)
	player := playerID2Player[playerID]

	if (player.appCoins < total) {
		a.WriteMsg(&msg.S2U_BuyResult{ErrorCode: 3})
		return
	}
	player.appCoins = player.appCoins - total
	playerID2Player[playerID] = player
	//购买操作
	for i := 0; i < count; i++ {
		//写节点
		block := WriteBlockMachine(playerID, productID, record.Price, 1)

		//写仓库数据
		AddMineMachine(block.Index, productID, playerID, 0)
	}
	log.Debug("handleU2S_Buy  productID:%v,count:%v", productID, count)

	//给发送者回应消息
	a.WriteMsg(&msg.S2U_BuyResult{ErrorCode: 0})

	//更新完成数据
	player.send()
	onShopList(a)
	onMineMachineList(a, playerID)
}

//矿机列表
func handleU2S_MineMachineList(args []interface{}) {

	//m := args[0].(*msg.U2S_MineMachineList)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}

	playerID := a.UserData().(uint)
	onMineMachineList(a, playerID)
}

func onMineMachineList(a gate.Agent, playerID uint) {
	sndmsg := &msg.S2U_MineMachineList{}
	list := GetMineMachineByPlayerID(playerID)
	timestamp := common.GetTimestamp()

	for _, r := range list {

		if r.Location > 0 {
			runTime := timestamp - r.StartTime
			outputCoin, _ := calGatherMine(r)
			p1 := msg.MineMachine{
				ID:         int32(r.ID),
				ProductID:  int32(r.DataID),
				Location:   int32(r.Location),
				RunTime:    runTime,
				EndTime:    r.EndTime,
				OutputCoin: float32(outputCoin),
			}
			sndmsg.List = append(sndmsg.List, &p1)
		} else {
			p1 := msg.MineMachine{
				ID:         int32(r.ID),
				ProductID:  int32(r.DataID),
				Location:   0,
				RunTime:    0,
				EndTime:    0,
				OutputCoin: 0.0,
			}
			sndmsg.List = append(sndmsg.List, &p1)
		}
	}
	a.WriteMsg(sndmsg)
}

//开始采矿
func handleU2S_Mining(args []interface{}) {
	m := args[0].(*msg.U2S_Mining)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}
	id := int(m.ID)
	playerID := a.UserData().(uint)
	location := int(m.Location)
	days := int(m.WillDay)

	if days < 1 {
		return
	}
	mineRecord, exist := gamedata.GetMineLocationByID(location)
	if !exist {
		return
	}
	if GetMineLocationFreeCount(location) < 1 {
		a.WriteMsg(&msg.S2U_MiningResult{ErrorCode: 1})
		return
	}

	record, exists := GetMineMachine(id)
	if !exists {
		a.WriteMsg(&msg.S2U_MiningResult{ErrorCode: 2})
		return
	}
	if record.PlayerID != playerID {
		a.WriteMsg(&msg.S2U_MiningResult{ErrorCode: 3})
		return
	}

	//扣费处理 （电费+管理费）*天数
	value := (mineRecord.Electric + mineRecord.ManageCost) * float32(days)
	player := playerID2Player[playerID]
	if player.appCoins < value {
		a.WriteMsg(&msg.S2U_MiningResult{ErrorCode: 4})
		return
	}

	player.appCoins = player.appCoins - value
	playerID2Player[playerID] = player
	player.onSaveDB()
	player.send()

	record.StartTime = common.GetTimestamp()
	record.EndTime = common.GetTimestamp() + int64(days*86400)
	record.Location = location
	SetMineMachine(id, record)

	a.WriteMsg(&msg.S2U_MiningResult{ErrorCode: 0})

	onMineMachineList(a, playerID)
}

//收集矿
func handleU2S_GatherMine(args []interface{}) {

	m := args[0].(*msg.U2S_GatherMine)
	// 消息的发送者
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}

	id := int(m.ID)
	//locate := m.Location
	playerID := a.UserData().(uint)

	onGatherMine(a, id, playerID)
}

func onGatherMine(a gate.Agent, id int, playerID uint) int32 {
	record, exists := GetMineMachine(id)
	if !exists {
		return 1
	}
	if record.PlayerID != playerID {
		return 2
	}

	//计算收益 todo
	value, freeMinute := calGatherMine(record)

	//改矿机时间
	restartTime := common.GetTimestamp() - freeMinute
	record.StartTime = restartTime
	SetMineMachine(id, record)

	//改帐号的数据

	a.WriteMsg(&msg.S2U_GatherMine{ID: int32(id), Btc: value})

	fmt.Println("Btc=", value)
	//更新消息
	player := playerID2Player[playerID]
	player.btc = player.btc + value
	playerID2Player[playerID] = player
	player.onSaveDB()
	player.send()

	//刷新仓库
	onMineMachineList(a, playerID)
	return 0
}

//计算收益
func calGatherMine(r MineMachine) (float32, int64) {
	if r.EndTime <= r.StartTime {
		return 0.0, 0
	}
	//收益公式：每天收益(1TH/s)*算力/24*已过小时数
	//计算收益, todo暂时定为1秒0.1个BTC
	record, exists := gamedata.GetShopByID(r.DataID)
	if !exists {
		return 0.0, 0
	}
	ghs := record.Ghs
	//为了测试方便，默认为1小时，正式使用前，请打开此处
	//useHours := (common.GetTimestamp() - r.StartTime) / 3600
	useHours := 1
	freeMinute := (common.GetTimestamp() - r.StartTime) % 3600
	//timeLen := float32(common.GetTimestamp() - r.StartTime)

	value := PeerDayBtc * ghs / 24 * float32(useHours)
	return value, freeMinute
}

//请求矿点信息
func handleU2S_MineList(args []interface{}) {
	//m := args[0].(*msg.U2S_MineMachineList)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}

	//矿点信息
	sndmsg1 := &msg.S2U_MineList{}
	minelist := gamedata.MineLocationData
	for _, r1 := range minelist {
		//todo 这里电费，管理费，产出率，可空余机位可调，待做
		p1 := &msg.MineInfo{
			Location: int32(r1.Location),
			Electric: r1.Electric,
			Tip:      r1.ManageCost,
			Ratio:    r1.Output,
			Num:      int32(GetMineLocationFreeCount(r1.Location)),
		}
		sndmsg1.List = append(sndmsg1.List, p1)
	}
	//fmt.Println("minelist=", sndmsg1)
	a.WriteMsg(sndmsg1)
}

//停止挖矿
func handleU2S_StopMining(args []interface{}) {
	m := args[0].(*msg.U2S_StopMining)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}

	id := int(m.ID)
	playerID := a.UserData().(uint)

	//验证数据
	record, exists := GetMineMachine(id)
	if !exists {
		a.WriteMsg(&msg.S2U_StopMiningResult{ErrorCode: 1})
		return
	}
	if record.PlayerID != playerID {
		a.WriteMsg(&msg.S2U_StopMiningResult{ErrorCode: 2})
		return
	}

	//给收益
	errorCode := onGatherMine(a, id, playerID)
	a.WriteMsg(&msg.S2U_StopMiningResult{ID: m.ID, ErrorCode: errorCode})

	//放进仓库
	record.StartTime = common.GetTimestamp()
	record.Location = 0
	SetMineMachine(id, record)

	//刷新仓库
	onMineMachineList(a, playerID)
}

//修改玩家信息
func handleU2S_ModifyPlayer(args []interface{}) {
	m := args[0].(*msg.U2S_ModifyPlayer)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}
	playerID := a.UserData().(uint)

	//名字为空，默认改的是头像
	if len(m.Name) > 0 {
		//检查名字是否重名
		if checkNameIsExits(m.Name) {
			a.WriteMsg(&msg.S2U_ModifyPlayer{ErrorCode: 1})
			return
		}
	}

	//更新消息
	player := playerID2Player[playerID]
	if len(m.Name) > 0 {
		player.name = m.Name
	}
	player.style = m.Style
	playerID2Player[playerID] = player
	player.onSaveDB()
	a.WriteMsg(&msg.S2U_ModifyPlayer{ErrorCode: 0})
	player.send()
}

//修改玩家信息
func handleU2S_Rank(args []interface{}) {
	//m := args[0].(*msg.U2S_Rank)
	a := args[1].(gate.Agent)
	if a.UserData() == nil {
		return
	}
	playerID := a.UserData().(uint)

	//最大名次
	var maxRow uint32
	maxRow = 10

	//资产排行榜
	sndmsg := &msg.S2U_AppCoinRank{}
	for _, r := range AppRank {
		if r.noID > maxRow {
			continue
		}

		p := &msg.RankInfo{
			NoID:     r.noID,
			PlayerID: uint32(r.playerID),
			Name:     r.name,
			Style:    r.style,
			Value:    r.value,
		}
		sndmsg.List = append(sndmsg.List, p)
	}
	//增加自己的排行榜数据
	r, exists := GetRankByPlayerID(playerID, 1)
	if exists {
		p2 := &msg.RankInfo{
			NoID:     r.noID,
			PlayerID: uint32(r.playerID),
			Name:     r.name,
			Style:    r.style,
			Value:    r.value,
		}
		sndmsg.List = append(sndmsg.List, p2)
	}
	a.WriteMsg(sndmsg)

	fmt.Println("S2U_AppCoinRank:",sndmsg)
	//今日收益排行榜
	sndmsg2 := &msg.S2U_BtcRank{}
	for _, r := range BtcRank {
		if r.noID > maxRow {
			continue
		}
		p1 := &msg.RankInfo{
			NoID:     r.noID,
			PlayerID: uint32(r.playerID),
			Name:     r.name,
			Style:    r.style,
			Value:    r.value,
		}
		sndmsg2.List = append(sndmsg2.List, p1)
	}

	//增加自己的排行榜数据
	r, exists = GetRankByPlayerID(playerID, 2)
	if exists {
		p := &msg.RankInfo{
			NoID:     r.noID,
			PlayerID: uint32(r.playerID),
			Name:     r.name,
			Style:    r.style,
			Value:    r.value,
		}
		sndmsg2.List = append(sndmsg2.List, p)
	}
	a.WriteMsg(sndmsg2)
	fmt.Println("sndmsg2:",sndmsg2)
}
