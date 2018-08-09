package internal

import (
	"server/mysql"
	"github.com/name5566/leaf/log"
	"server/gamedata"
)

//玩家矿机信息
type MineMachine struct {
	ID        int `gorm:"primary_key"` //矿机实例ID
	DataID    int                      //矿机配置ID
	PlayerID  uint                     //玩家ID
	Location  int                      //矿机存放位置，0：仓库，其它：矿场点编号
	StartTime int64                    //挖矿开始时间
	EndTime   int64                    //挖矿结束时间
}

var (
	MineMachineList = make(map[int]MineMachine)
)

func MineMachineInit() {
	//加载数据库
	db := mysql.MysqlDB()
	rows, err := db.Query("select `id`,`data_id`, `player_id`, `location`, `start_time`,`end_time` from `mine_machines`" +
		" where `location` <>0 order by `start_time` desc limit 8192")
	if nil == err {
		for rows.Next() {
			u := MineMachine{}
			err = rows.Scan(&u.ID, &u.DataID, &u.PlayerID, &u.Location, &u.StartTime, &u.EndTime)
			if nil == err {
				MineMachineList[u.ID] = u
			}
		}
	}
	if nil != err {
		log.Error("MineMachineInit error:", err)
	}
}

func MineMachineLoad(playerID uint) {
	//加载数据库
	db := mysql.MysqlDB()
	stmt, err := db.Prepare("select `id`,`data_id`, `player_id`, `location`, `start_time`,`end_time` from `mine_machines` WHERE `player_id`=?")
	if nil == err {
		rows, err := stmt.Query(playerID)
		if nil == err {
			for rows.Next() {
				u := MineMachine{}
				err = rows.Scan(&u.ID, &u.DataID, &u.PlayerID, &u.Location, &u.StartTime, &u.EndTime)
				if nil == err {
					MineMachineList[u.ID] = u
				}
			}
		}
	}

	if nil != err {
		log.Error("MineMachineInit error:", err)
	}
}

func GetMineMachine(id int) (MineMachine, bool) {
	record, exists := MineMachineList[id]
	return record, exists
}

func GetMineMachineByPlayerID(playerID uint) ([]MineMachine) {

	var list []MineMachine
	for _, record := range MineMachineList {
		if (record.PlayerID == playerID) {
			list = append(list, record)
		}
	}
	return list
}

func AddMineMachine(objectID int, dataID int, fromPlayerID uint, fromLocation int) bool {
	_, exists := MineMachineList[objectID]
	if (exists) {
		return false
	}

	p1 := &MineMachine{
		ID:       objectID,
		DataID:   dataID,
		PlayerID: fromPlayerID,
		Location: fromLocation,
	}
	MineMachineList[objectID] = *p1

	db := mysql.MysqlDB()
	stmt, err := db.Prepare("insert mine_machines values(?,?,?,?,?,?)")
	if nil == err {
		_, err = stmt.Exec(p1.ID, p1.DataID, p1.PlayerID, p1.Location, p1.StartTime, p1.EndTime)
	}
	if nil != err {
		log.Error("create AddMineMachine error:", err)
	}

	return true
}

func SetMineMachine(objectID int, record MineMachine) {
	MineMachineList[objectID] = record

	db := mysql.MysqlDB()
	stmt, err := db.Prepare("update mine_machines set location=?,start_time=?,end_time=? where id =?")
	if nil == err {
		_, err = stmt.Exec(record.Location, record.StartTime, record.EndTime, record.ID)
	}
	if nil != err {
		log.Error("SetMineMachine error:", err)
	}
}

//清除玩家所属矿机
func DeleteMineMachine(objectID int) {
	delete(MineMachineList, objectID)

	db := mysql.MysqlDB()
	stmt, err := db.Prepare("delete from table mine_machines where id =?")
	if nil == err {
		_, err = stmt.Exec(objectID)
	}

	if nil != err {
		log.Error("DeleteMineMachine error:", err)
	}
}

//得到某个矿点的空闲机位数
func GetMineLocationFreeCount(location int) int {
	value := 0
	record, exists := gamedata.GetMineLocationByID(location)
	if !exists {
		return 0
	}

	for _, r := range MineMachineList {
		if (r.Location == location) {
			value ++
		}
	}
	if record.Num <= value {
		return 0
	}
	return record.Num - value

}
