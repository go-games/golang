package gate

import (
	"server/msg"
	"server/msg/protocol"
	"server/login"
)

func init() {
     msg.Processor.SetRouter(&protocol.C_Register{},login.ChanRPC)
     msg.Processor.SetRouter(&protocol.C_LoginByPwd{},login.ChanRPC)
}
