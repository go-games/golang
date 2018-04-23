package gate

import (
	"server/msg"
	"server/login"
)

func init() {
	msg.Processor.SetRouter(&msg.UserLogin{},login.ChanRPC)
}
