package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"fmt"
	"github.com/name5566/leaf/log"
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
	InitGameTables()
	OnTickHour()

	log.Release("gameModule Ready ok")
	fmt.Println("gameModule Ready ok")
}

func InitGameTables() {
	UpdatePeerDayBtc()
	MineMachineInit()
	DepotChainInit()
	RankInit()
}

func (m *Module) OnDestroy() {

}
