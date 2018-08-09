package main
import (
	_"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func main() {

	Test()
}


//type Player struct {
//	ID    uint
//	//初始化之后自己写函数读取BaseInfo和CardsInfo吧，保存信息也一样。
//	//本来想用ORM的级联发现各种不习惯。还是自己写代码控制下逻辑好了，也多不了几行代码。
//	playerBaseInfo PlayerBaseInfo
//	cards_map map[uint]PlayersCards // key is cardID , value is PlayersCards info
//
//	//...其他信息
//}

type PlayerBaseInfo struct {
	PlayerID    uint `gorm:"primary_key"`
	Name string `gorm:"not null;unique"`
	Style uint`gorm:"default:'0'"`
	AppCoins uint`gorm:"default:'0'"`
	ComputePower uint`gorm:"default:'0'"`
}

/*
type PlayersCards struct {
	ID uint
	PlayerID uint `gorm:"unique_index:idx_name_code"`
	CardID uint `gorm:"unique_index:idx_name_code"`
	Amount uint`gorm:"default:'0'"`
	Level uint `gorm:"default:'1'"`
}

type Card struct {
	ID uint `gorm:"primary_key"`
	Name string `gorm:"not null"` // e.a. Knights of the Round Table
}
*/

type Account struct {
	ID             uint `gorm:"primary_key"`
	Account        string `gorm:"not null;unique"`
	SdkAccount     string `protobuf:"bytes,1,opt,name=SdkAccount,proto3" json:"SdkAccount,omitempty"`
	SdkClientID    string `protobuf:"bytes,2,opt,name=SdkClientID,proto3" json:"SdkClientID,omitempty"`
	SdkAccessToken string `protobuf:"bytes,3,opt,name=SdkAccessToken,proto3" json:"SdkAccessToken,omitempty"`
	ChannelID      uint32 `protobuf:"varint,4,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	Imei           string `protobuf:"bytes,6,opt,name=Imei,proto3" json:"Imei,omitempty"`
	GameVer        uint32 `protobuf:"varint,11,opt,name=Game_ver,json=GameVer,proto3" json:"Game_ver,omitempty"`
	Extra          string `protobuf:"bytes,13,opt,name=Extra,proto3" json:"Extra,omitempty"`

	Phone string `gorm:"default:'0'`
	Createtime        int64 `gorm:"default:'0'"`

	Name         string `gorm:"not null;unique"`
	Style        uint32 `gorm:"default:'0'"`
	AppCoins     float32 `gorm:"default:'0'"`
	ComputePower float32 `gorm:"default:'0'"`
	Logintime    int64	//登录时间
	Offflinetime int64	//离线时间
	Btc 		 float32 `gorm:"default:'0'"`
}

type MineMachine struct {
	ID       int `gorm:"primary_key"` //矿机实例ID
	DataID   int						//矿机配置ID
	PlayerID uint                     //玩家ID
	Location int                     //矿机存放位置，0：仓库，其它：矿场点编号
	StartTime int64						//挖矿开始时间
	EndTime int64						//挖矿结束时间
}

type BlockMachine struct {
	Index     int
	Timestamp string
	PlayerID  uint
	DataID    int
	Price	  float32		//单价
	Count	  int			//数量
	Hash      string
	PrevHash  string
}

func Test() {
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/blockchain?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("connect db error")
	}
	defer db.Close()
	db.DropTableIfExists(&Account{},&MineMachine{},&BlockMachine{})
	//&Card{}, &PlayersCards{},&PlayerBase{})
	////
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&MineMachine{})
	db.AutoMigrate(&BlockMachine{})
	//db.AutoMigrate(&Card{})
	//db.AutoMigrate(&PlayersCards{})
	/*
		block := BlockMachine{PlayerID:1}
		err =  db.Save(&block).Error

		block = BlockMachine{PlayerID:2}
		err =  db.Save(&block).Err
		/*
		playersCard1 := PlayersCards{PlayerID:1,CardID:1,Amount:20}
		err =  db.Save(&playersCard1).Error
		if nil != err {
			fmt.Println("create 1 error:",err)
		}

		playersCard2 := PlayersCards{PlayerID:1,CardID:2,Amount:65}
		err =  db.Save(&playersCard2).Error
		if nil != err {
			fmt.Println("create 2 error:",err)
		}

		playersCard1.Amount = 0
		err =  db.Save(playersCard1).Error
		if nil != err {
			fmt.Println("Update 1 error:",err)
		}

		playersCard3 := PlayersCards{PlayerID:2,CardID:3,Amount:615}
		err =  db.Save(&playersCard3).Error
		if nil != err {
			fmt.Println("create 3 error:",err)
		}

		player1 := PlayerBaseInfo{ID:2,Name:"Mike11221"}
		err =  db.Save(&player1).Error
		if nil != err {
			fmt.Println("create 3 error:",err)
		}


		//err2 := db.Create(&Player{
		//	Name: "Mike",
		//}).Error
		//if nil != err2 {
		//	fmt.Println("already exist:",err2)
		//}


		//query player Mike
		var player PlayerBaseInfo
		err = db.Where("Name = ?", "Mike11221"). Limit(1).Find(&player).Error
		if nil != err {
			fmt.Println(err)
		}
		fmt.Println(player)
		*/

	type blockMachine struct {
		Index     int
		Timestamp string
		PlayerID  uint
		DataID    int
		Price     int //单价
		Count     int //数量
		Hash      string
		PrevHash  string
	}

	//读矿机连数据
	var list []blockMachine

	db.Table("block_machines").Select("*").Scan(&list)
	//rows, err := db.Table("block_machines").Select("*").Rows()
	for _, r := range list {
		println(r.PlayerID,r.Index)
	}
	println(list)
}

