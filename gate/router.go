package gate

import (
	"golang/msg"
	"golang/game"
)

func init() {
	msg.Processor.SetRouter(&msg.UserLogin{}, game.ChanRPC)
}
