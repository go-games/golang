package internal

import (
	"fmt"
	"reflect"
	"time"

	"server/msg"

	"github.com/name5566/leaf/db/dbos"
	"github.com/name5566/leaf/gate"
)

func init() {
	handler(&msg.HelloWorld, handleHello)
}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	fmt.Println("hello %v", m.Name)
	dbos.ExistPlayer(m.Name)
	// 给发送者回应一个 Hello 消息

	a.WriteMsg(&msg.Hello{
		Name: "122132",
	})
	time.Sleep(10 * time.Second)

}
