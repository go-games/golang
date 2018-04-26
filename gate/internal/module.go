package internal

import (
	"github.com/name5566/leaf/gate"
	"golang/conf"
	"golang/game"
	"golang/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	config:=conf.GetInstance()
	m.Gate = &gate.Gate{
		MaxConnNum:      config.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          config.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        config.Server.CertFile,
		KeyFile:         config.Server.KeyFile,
		TCPAddr:         config.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
