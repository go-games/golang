package gate

import (
	"golang/msg"
	"golang/game"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
