package internal

import (
	"reflect"
	"server/db"

	 p "server/msg/protocol"
	"github.com/name5566/leaf/gate"
	"strconv"
	"time"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&p.C_Register{},RegisterPlayer)
	handleMsg(&p.C_LoginByPwd{},PlayerLogin)
}
func PlayerLogin(args []interface{}){
	m := args[0].(*p.C_LoginByPwd)
	a := args[1].(gate.Agent)
	if CheckPlayer(m.UserName,m.Pwd) {
	//有待改进
		a.WriteMsg(&p.LoginSuccess{
			Error:"",
			Code:200,
			Result: struct {
				UserId, UserName string
				IsRobt, IsGuest  bool
				State            int
			}{UserId:Rand_User_id(),UserName:m.UserName,IsRobt:false,IsGuest:false , State: 1},
		})
	}else {
		a.WriteMsg(&p.Failed{Error:"用户名或密码错误",Code:2003})
		return
	}
}
func RegisterPlayer(args []interface{}){
	m := args[0].(*p.C_Register)
	a := args[1].(gate.Agent)
	if ExistPlayer(m.UserName) {
		a.WriteMsg(&p.Failed{"该用户已存在",2003,""})
		return
	}
	err := db.Insert(m)
	if err != nil {
		a.WriteMsg(&p.Failed{err.Error(),2003,""})
		return
	}
	//该写法有待改进
	a.WriteMsg(&p.RegisterSuccess{
		Error:"",
		Code:200,
		Result: struct {
			UserId, UserName string
			IsRobt, IsGuest  bool
			State            int
		}{UserId:Rand_User_id(), UserName:m.UserName, IsRobt: false, IsGuest: false, State:1 },

	})
}
func ExistPlayer(username string) bool{
		var result interface{}
		type condition struct {
			UserName string
		}
		c := condition{UserName:username}
	  	err := db.FindOnce(c,&result)
	    log.Debug("%v",result)
		if err != nil && result == nil{
			return false
		}
		return true
	}
func CheckPlayer(username,password string) bool {
	var result interface{}
	type condition struct {
		UserName string
		Pwd string
	}
	c := condition{UserName:username,Pwd:password}
	err := db.FindOnce(c,&result)
	if err != nil {
		return false
	}
	return true
}
func Rand_User_id() string {
	return strconv.FormatInt(time.Now().UnixNano(),10)
}