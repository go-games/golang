package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

type UserLogin struct {
	UerId string `json:"LoginName"`
	Passwd string `json:"LoginPw"`
}
type Resp struct{
	RetCode int
	RetMsg  string
	RespData interface{}
}
func init() {
	Processor.Register(&UserLogin{})
	Processor.Register(&Resp{})
}
