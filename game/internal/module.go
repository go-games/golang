package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"server/fight"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	fight.Run()
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
