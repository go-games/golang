package msg

import (
	"github.com/name5566/leaf/network/json"
	"server/fight"
	"server/utils"
	p "server/msg/protocol"//协议，建议放一起
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Hello{})
	Processor.Register(&Roomresp{})
	Processor.Register(&FightResp{})
	Processor.Register(&p.C_Register{})
	Processor.Register(&p.C_LoginByPwd{})
	Processor.Register(&p.RegisterSuccess{})
	Processor.Register(&p.LoginSuccess{})
	Processor.Register(&p.Failed{})
	Processor.Register(&utils.Result{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
	Name string
}

type Roomresp struct {
	RoomId   string
	ServerId string
	UserId   string
	ProtoId  int //协议ID 根据id来区分
}

type FightResp struct {
	RoomId    string
	ServerId  string
	WinUserId int
	Data      fight.Keyframe
	ProtoId   int //协议ID 根据id来区分
}
