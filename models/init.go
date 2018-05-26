package models


import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
)


var c *mongodb.DialContext

func init() {
	var err error
	c, err = mongodb.Dial("mongodb://game:golang.ltd@114.55.73.227:7770/game", 30)
	if err != nil {
		log.Fatal("mongodb连接失败")
		return
	}



}


//初始化 对外部暴露
var RoomService = &roomService{}


//TODO 这里需要注意下，临时先写上
func MogClose() {
	c.Close()
}