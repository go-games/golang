package gate

import (
	"server/game"
	"server/msg"
	"server/msg/protocol"
	"server/login"
)

func init() {
	//msg.Processor.SetRouter(&msg.Room_resp{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Roomresp{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.FightResp{}, game.ChanRPC)

	msg.Processor.SetRouter(&protocol.C_Register{},login.ChanRPC)
	msg.Processor.SetRouter(&protocol.C_LoginByPwd{},login.ChanRPC)
	//room.Processor.SetRouter(&room.Create{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Room_enter{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Room_quit{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Room_quick_enter{}, game.ChanRPC)

}
