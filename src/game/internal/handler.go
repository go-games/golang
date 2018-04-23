package internal

import (
	"reflect"
)
func init(){
	//handlerMsg(&msg.Hello{},HandleHello)
}
func handlerMsg(m interface{},h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)

}
func HandleHello(args [] interface{}) {
}
func handleUser(agrs [] interface{}) {

}