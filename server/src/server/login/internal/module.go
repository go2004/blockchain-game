package internal

import (

	"github.com/name5566/leaf/module"
	"server/base"
	"github.com/name5566/leaf/log"
	"fmt"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}



func (m *Module) OnInit() {
	m.Skeleton = skeleton
	InitLoginTables()
	log.Release("loginModule Ready ok")
	fmt.Println("loginModule Ready ok")
}

func InitLoginTables() {
	//db := mysql.MysqlDB()
	//db.AutoMigrate(&Account{})
}


func (m *Module) OnDestroy() {

}
