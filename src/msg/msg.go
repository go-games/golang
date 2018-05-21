package msg

import (
	"github.com/name5566/leaf/network/json"
	 p "server/msg/protocol"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&p.C_Register{})
	Processor.Register(&p.C_LoginByPwd{})
	Processor.Register(&p.RegisterSuccess{})
	Processor.Register(&p.LoginSuccess{})
	Processor.Register(&p.Failed{})
}
