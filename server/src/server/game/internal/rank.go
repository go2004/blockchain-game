package internal

import (
	"server/mysql"
	"github.com/name5566/leaf/log"
)

//var (
//	AppRank = make(map[uint]*rankInfo)
//	BtcRank = make(map[uint]*rankInfo)
//)

//%%游戏币排名
type RankInfo struct {
	//角色信息
	playerID uint    //玩家ID
	name     string  //玩家名字
	style    uint32  //外形头像
	value    float32 //数据
	noID     uint32  //排名
}

var AppRank []RankInfo
var BtcRank []RankInfo

//从数据库加载
func RankInit() {

	//游戏币
	db := mysql.MysqlDB()
	rows, err := db.Query("SELECT id,`name`,style,app_coins FROM accounts ORDER BY app_coins desc")
	//rows, err := db.Query("SELECT id,`name`,style,app_coins,(SELECT count(DISTINCT app_coins) " +
	//	"FROM accounts AS b WHERE a.app_coins<b.app_coins)+1 AS rank	FROM accounts AS a ORDER BY app_coins")
	if nil == err {
		var noID uint32
		noID = 1
		for rows.Next() {
			u := RankInfo{}
			err = rows.Scan(&u.playerID, &u.name, &u.style, &u.value)
			if nil == err {
				u.noID = noID
				AppRank = append(AppRank, u)
				noID++
			}
		}
	}

	//比特币
	rows, err = db.Query("SELECT id,`name`,style,btc FROM accounts ORDER BY btc desc")
	if nil == err {
		var noID uint32
		noID = 1
		for rows.Next() {
			u := RankInfo{}
			err = rows.Scan(&u.playerID, &u.name, &u.style, &u.value)
			if nil == err {
				u.noID = noID
				BtcRank = append(BtcRank, u)
				noID++
			}
		}
	}
	if nil != err {
		log.Error("RankInit error:", err)
	}
}

func GetRank(ranktype int) []RankInfo {

	return nil
}

func GetRankByPlayerID(playerId uint, rankType int) (RankInfo, bool) {
	var r1 RankInfo
	rankList := AppRank
	if rankType == 2 {
		rankList = BtcRank
	}
	for _, r := range rankList {
		if r.playerID == playerId {
			return r, true
		}
	}
	return r1, false
}
