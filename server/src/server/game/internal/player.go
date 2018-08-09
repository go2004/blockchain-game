package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/timer"
	"github.com/name5566/leaf/go"
	"time"
	"server/msg"
	"server/mysql"
	"github.com/name5566/leaf/log"
	"fmt"
)

var (
	playerID2Player = make(map[uint]*Player)
)

const (
	userLogin  = iota
	userLogout
	userGame
)

type Player struct {
	gate.Agent
	*g.LinearContext
	state       int
	saveDBTimer *timer.Timer

	//帐号信息
	account   string
	channelID uint32
	playerID  uint

	//角色信息
	name         string
	style        uint32
	appCoins     float32
	computePower float32
	logintime    int64
	offflinetime int64
	btc          float32

}

func (player *Player) login(playerID uint) {
	skeleton.Go(func() {
		MineMachineLoad(playerID)
	}, func() {
		// network closed
		if player.state == userLogout {
			player.logout(playerID)
			return
		}

		// db error
		player.state = userGame
		playerID2Player[playerID] = player

		player.onLogin()
		player.autoSaveDB()
	})

	player.WriteMsg(&msg.S2U_LoginResult{Result: 0})
	//log.Debug("%v:S2U_LoginResult{Result: 0} ", playerID)
	player.send()
}
//更新玩家数据
func (player *Player) send() {
	sndMsg := msg.S2U_PlayerBase{
		PlayerID:     uint32(player.playerID),
		Account:      player.account,
		Name:         player.name,
		Style:        player.style,
		AppCoins:     player.appCoins,
		ComputePower: player.computePower,
		Btc:          player.btc,
	}
	player.WriteMsg(&sndMsg)
	fmt.Println("S2U_PlayerBase ", sndMsg)
}

func (player *Player) isOffline() bool {
	return player.state == userLogout
}

func (player *Player) logout(playerID uint) {
	player.onLogout()
}

func (player *Player) autoSaveDB() {
	const duration = 5 * time.Minute
	// save
	player.saveDBTimer = skeleton.AfterFunc(duration, func() {

		player.Go(func() {
			player.onSaveDB()

		}, func() {
			player.autoSaveDB()
		})
	})
}

func (player *Player) onLogin() {

}

func (player *Player) onLogout() {
	log.Debug("player Logout %v", player.playerID)
	player.onSaveDB()
}

func (player *Player) onSaveDB() {
	db := mysql.MysqlDB()
	db.Exec("update accounts set name=?,style=?,app_coins=?,compute_power=?,btc=? where id =?",
		player.name, player.style, player.appCoins, player.computePower, player.btc, player.playerID)
}

//重名检查
func checkNameIsExits(NewName string) bool {
	db := mysql.MysqlDB()
	stmt, err := db.Prepare("select `id` from `accounts` WHERE `name`=? limit 1")
	accountid := 0
	if nil == err {
		rows, err := stmt.Query(NewName)
		if nil == err {
			for rows.Next() {
				err = rows.Scan(&accountid)
			}
		}
	}
	if nil != err {
		log.Error("checkNameIsExits error:", err)
		return true
	}

	return accountid > 0
}
