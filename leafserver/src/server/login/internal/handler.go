package internal

import (
	"reflect"
	"server/msg"
	"time"

	"github.com/name5566/leaf/db/dbos"
	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.UserLogin{}, handleTest)
}
func handleTest(args []interface{}) {
	// 收到的 Test 消息
	unixtime := time.Now().Unix()
	m := args[0].(*msg.UserLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 1 查询数据库--判断用户是不是合法
	if exist, err := dbos.ExistPlayer(m.LoginName); err != nil {
		//TODO错误处理
		return
	} else if !exist {
		//TODO错误处理
		return
	}
	if err := dbos.InsertPlayer(m.LoginName, m.LoginPW, "test", unixtime); err != nil {
		//TODO错误处理
		return
	}
	// 2 如果数据库返回查询正确--保存到缓存或者内存
	//TODO redis 讨论如何缓存

	a.WriteMsg(&msg.UserLogin{
		LoginName: "client",
	})
}
