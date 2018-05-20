package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	//msg.Processor.SetRouter(&msg.Room_resp{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Roomresp{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Create{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Room_enter{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Room_quit{}, game.ChanRPC)
	//room.Processor.SetRouter(&room.Room_quick_enter{}, game.ChanRPC)

}
