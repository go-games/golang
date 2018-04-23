package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"logic"
	"fmt"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}


func init() {
	handleMsg(&msg.UserLogin{},handleLogin)

}
func handleLogin(args []interface{}){
	m := args[0].(*msg.UserLogin)
	a := args[1].(gate.Agent)
	//defer a.Close()
	status,err := logic.CheckUser(m.UerId,m.Passwd)
	if !status {
		a.WriteMsg(&msg.Resp{403,fmt.Sprint("%s",err),""})
		return
	}else {
		a.WriteMsg(&msg.Resp{200,"登陆成功",""})
	}

}
