package login

import (
	"server/login/internal"
	"github.com/name5566/leaf/gate"
	"sync"
	"server/msg/protocol"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
//UserId
var UserSession sync.Map
func CheckLogin(key string)bool{
	var a interface{}
	_,ok := UserSession.Load(key)
	if ok {
		return true
	}
	a.(gate.Agent).WriteMsg(&protocol.Failed{Error:"还未登陆",Code:4003,Result:""})
	return false
}
