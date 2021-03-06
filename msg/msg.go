package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&UserLogin{})
}


// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type UserLogin struct {
	LoginName string
	LoginPW   string
}