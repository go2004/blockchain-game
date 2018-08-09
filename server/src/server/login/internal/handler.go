package internal

import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/log"
	"server/common"
	_ "strconv"
	"strings"
	"server/game"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.U2S_Login{}, handleAuth)
	handleMsg(&msg.U2S_RegistAccount{}, handleRegistAccount)
}
func handleAuth(args []interface{}) {
	m := args[0].(*msg.U2S_Login)
	a := args[1].(gate.Agent)

	//帐号长度异常
	if len(m.SdkAccount) < 2 || len(m.SdkAccount) > 12 {
		log.Debug("%v:S2U_LoginResult{Result: 1} ", m.SdkAccount)
		a.WriteMsg(&msg.S2U_LoginResult{Result: 1})
		return
	}

	accountID := makeAccount(m.SdkAccount, m.ChannelID)
	accountInfo := getAccountByAccountID(accountID)
	//data := []byte(m.SdkAccessToken)
	//var hash = md5.Sum(data)
	//password := hex.EncodeToString(hash[:])
	// 帐号不存在
	if nil == accountInfo {
		a.WriteMsg(&msg.S2U_LoginResult{Result: 4})
		log.Debug("%v:S2U_LoginResult{Result: 4} ", m.SdkAccount)
		return
	}

	// match password
	if m.SdkAccessToken != accountInfo.SdkAccessToken {
		a.WriteMsg(&msg.S2U_LoginResult{Result: 1})
		log.Debug("%v:S2U_LoginResult{Result: 1} ", m.SdkAccount)
		return

	}

	a.SetUserData(accountInfo.ID)
	game.ChanRPC.Go("UserLogin", a, accountInfo.ID, accountInfo.Account, accountInfo.ChannelID, accountInfo.Name, accountInfo.Style,
		accountInfo.AppCoins,accountInfo.ComputePower, accountInfo.Btc)
	log.Debug("%v:S2U_LoginResult{Result: 0} ", m.SdkAccount)
	//fmt.Println("%v:S2U_LoginResult{Result: 0} ", m.SdkAccount)
}

//生成帐号 帐号生成规则，SDK帐号+KEY+渠道编号
func makeAccount(sdkAccount string, channelID int32) string {
	accountKey := "0Vf363e6Gd0yz41m" //游戏服的帐名key
	sdkAccount1 := strings.Replace(sdkAccount, "-", "*", 0)
	accountStr := sdkAccount1 + "-" + accountKey + "-" + string(channelID)
	newAccount := common.MD5(accountStr)
	return newAccount
}

//创建帐号信息
func handleRegistAccount(args []interface{}) {
	m := args[0].(*msg.U2S_RegistAccount)
	a := args[1].(gate.Agent)

	//处理过滤字 todo

	//帐号长度异常
	if len(m.SdkAccount) < 2 || len(m.SdkAccount) > 12 {
		log.Debug("%v:G2U_RegistAccountResult{ErrorCode: 1} ", m.SdkAccount)
		a.WriteMsg(&msg.G2U_RegistAccountResult{ErrorCode: 1})
		return
	}

	accountID := makeAccount(m.SdkAccount, m.ChannelID)

	accountInfo := getAccountByAccountID(accountID)
	if nil != accountInfo {
		log.Debug("%v:G2U_RegistAccountResult{ErrorCode: 5} ", m.SdkAccount)
		a.WriteMsg(&msg.G2U_RegistAccountResult{ErrorCode: 5})
		return
	}

	//not having this account,creat account
	newAccount := RegistAccount(accountID, m)
	if nil != newAccount {
		log.Debug("%v:G2U_RegistAccountResult{ErrorCode: 0} ", m.SdkAccount)
		a.WriteMsg(&msg.G2U_RegistAccountResult{ErrorCode: 0})
		a.SetUserData(newAccount.ID)
		game.ChanRPC.Go("UserLogin", a, newAccount.ID, newAccount.Account, newAccount.ChannelID, newAccount.Name, newAccount.Style,
			newAccount.AppCoins, newAccount.ComputePower,newAccount.Btc)

	} else {
		log.Release("%v:G2U_RegistAccountResult{ErrorCode: 2} ", m.SdkAccount)
		a.WriteMsg(&msg.G2U_RegistAccountResult{ErrorCode: 2})
	}

}
