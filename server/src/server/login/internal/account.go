package internal

import (
	"server/mysql"
	"server/msg"
	"server/common"
	"github.com/name5566/leaf/log"
)

type Account struct {
	ID             uint   `gorm:"primary_key"`
	Account        string `gorm:"not null;unique"`
	SdkAccount     string `protobuf:"bytes,1,opt,name=SdkAccount,proto3" json:"SdkAccount,omitempty"`
	SdkClientID    string `protobuf:"bytes,2,opt,name=SdkClientID,proto3" json:"SdkClientID,omitempty"`
	SdkAccessToken string `protobuf:"bytes,3,opt,name=SdkAccessToken,proto3" json:"SdkAccessToken,omitempty"`
	ChannelID      uint32 `protobuf:"varint,4,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	Imei           string `protobuf:"bytes,6,opt,name=Imei,proto3" json:"Imei,omitempty"`
	GameVer        uint32 `protobuf:"varint,11,opt,name=Game_ver,json=GameVer,proto3" json:"Game_ver,omitempty"`
	Extra          string `protobuf:"bytes,13,opt,name=Extra,proto3" json:"Extra,omitempty"`

	Phone      string `gorm:"default:'0'`
	Createtime int64  `gorm:"default:'0'"`

	Name         string  `gorm:"not null;unique"`
	Style        uint32  `gorm:"default:'0'"`
	AppCoins     float32 `gorm:"default:'0'"`
	ComputePower float32 `gorm:"default:'0'"`
	Logintime    int64 //登录时间
	Offflinetime int64 //离线时间
	Btc          float32 `gorm:"default:'0'"`
}

func getAccountByAccountID(accountID string) *Account {

	var account Account
	db := mysql.MysqlDB()
	stmt, err := db.Prepare("SELECT * FROM `accounts` WHERE `account` =?")
	if nil == err {
		rows, err1 := stmt.Query(accountID)
		if nil == err1 {
			for rows.Next() {
				err1 = rows.Scan(&account.ID, &account.Account, &account.SdkAccount, &account.SdkClientID, &account.SdkAccessToken,
					&account.ChannelID, &account.Imei, &account.GameVer, &account.Extra, &account.Phone,
					&account.Createtime, &account.Name, &account.Style, &account.AppCoins, &account.ComputePower,
					&account.Logintime, &account.Offflinetime, &account.Btc)
			}
		}
		err = err1
	}

	if nil != err {
		log.Error("getAccountByAccountID error:", err)
	}
	if 0 == account.ID {
		return nil
	} else {
		return &account
	}

}

func RegistAccount(accountID string, m *msg.U2S_RegistAccount) *Account {
	var p = Account{
		Account:        accountID,
		SdkAccount:     m.SdkAccount,
		SdkClientID:    m.SdkClientID,
		SdkAccessToken: m.SdkAccessToken,
		ChannelID:      uint32(m.ChannelID),
		Imei:           m.Imei,
		Name:           m.Name,
		Style:          m.Style,
		Phone:          m.Phone,
		AppCoins:       1000000.0,
		Createtime:     common.GetTimestamp(),
	}

	db := mysql.MysqlDB()

	sql := "insert accounts(`account`,`sdk_account`,`sdk_client_id`,`sdk_access_token`,`channel_id`," +
		"`imei`,`game_ver`,`extra`,`phone`,`createtime`," +
		"`name`,`style`,`app_coins`,`compute_power`,`logintime`," +
		"`offflinetime`,`btc`) " +
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if nil == err {
		res, err := stmt.Exec(p.Account, p.SdkAccount, p.SdkClientID, p.SdkAccessToken, p.ChannelID,
			p.Imei, p.GameVer, p.Extra, p.Phone, p.Createtime,
			p.Name, p.Style, p.AppCoins, p.ComputePower, p.Logintime,
			p.Offflinetime, p.Btc)
		if nil == err {
			id, err := res.LastInsertId()
			if nil == err {
				p.ID = uint(id)
			}
			if nil != err {
				log.Error("RegistAccount error:", err)
			}
		}
		if nil != err {
			log.Error("RegistAccount error:", err)
		}
	}
	if nil != err {
		log.Error("RegistAccount error:", err)
	}

	if 0 == p.ID {
		return nil
	} else
	{
		return &p
	}
}
