package chain

type mineMachine struct {
	ID       int `gorm:"primary_key"` //矿机实例ID
	DataID   int						//矿机配置ID
	PlayerID uint                     //玩家ID
	Location uint                     //矿机存放位置，0：仓库，其它：矿场点编号
}

var (
	MineMachineList = make(map[int]mineMachine)
)

func MineMachineInit() {
	//加载数据库
}

func GetMineMachine(ID int) (mineMachine, bool) {
	record, exists := MineMachineList[ID]
	return record, exists
}

func GetMineMachineByPlayerID(PlayerID uint) ([]mineMachine) {

	var list []mineMachine
	for _, record := range MineMachineList {
		if (record.PlayerID == PlayerID) {
			list = append(list, record)
		}
	}
	return list
}

func AddMineMachine(objectID int, dataID int, fromPlayerID uint, fromLocation uint) bool {
	_, exists := MineMachineList[objectID]
	if (exists) {
		return false
	}

	p1 := &mineMachine{
		ID:       objectID,
		DataID: 	dataID,
		PlayerID: fromPlayerID,
		Location: fromLocation,
	}
	MineMachineList[objectID] = *p1
	return true
}

//可以这里有风险，todo
func SetMineMachine(objectID int, record mineMachine) {
	MineMachineList[objectID] = record
}

func DeleteMineMachine(objectID int) {
	delete(MineMachineList, objectID)
}
