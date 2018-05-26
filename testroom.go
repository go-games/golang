package main

import (
	"server/room"
	"glog/glog"
	"flag"
)

func main() {
	flag.Parse()

	var a room.Room
	a =room.NewRoom()
	a.Get(2000)

	var user_right_id = "101"
	var user_left_id = "102"

	roominfo :=room.NewRoomInfo()
	roomid,state:=roominfo.Create(user_left_id)
	if !state {
		glog.Info("创建房间失败")
		return
	}
	roominfo.Join(user_right_id,roomid)

	roominfo.Join(user_left_id,roomid)
	roominfo.Join("103",roomid)
	people_num,state := roominfo.Quit(user_left_id)
	if state && people_num ==0 {
		roominfo.Close()
	}
	people_num,state = roominfo.Quit(user_right_id)
	if state && people_num ==0 {
		roominfo.Close()
	}


}
