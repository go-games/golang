package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"server/msg"
	"server/room"
	"server/utils"
	"server/conf"
	"fmt"
)


/*
	Name    string
	RoomId   string `json:"RoomId"`
	ServerId string `json:"ServerId"`
	UserId string `json:"UserId"`
	ProtoId  string `json:"ProtoId"`  //协议ID 根据id来区分
*/

func handleRoomresp(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Roomresp)
	// 消息的发送者
	a := args[1].(gate.Agent)

	switch m.ProtoId {
	case 1000:
		//创建房间
		code, resp, err:=room_create(m)
		a.WriteMsg(utils.Resp_result(code,resp,err))

	case 1001:
		//加入房间
		code, resp, err:=room_join(m)
		a.WriteMsg(utils.Resp_result(code,resp,err))

	case 1002:
		//退出房间
		code, resp, err:=room_quit(m)
		a.WriteMsg(utils.Resp_result(code,resp,err))

	case 1003:
		//随机加入一个房间
		code, resp, err:=room_quick_enter(m)
		a.WriteMsg(utils.Resp_result(code,resp,err))

	default:
		log.Error("room协议id不匹配",m.ProtoId)
		a.WriteMsg(utils.ErrGame("protoid not found","",404))
	}
}



func room_create(m *msg.Roomresp) (int ,*room.Room_create, string){
	roominfo :=room.NewRoomInfo()
	//创建房间，房间创建者自动加入放假
	roomid :=roominfo.Create(m.UserId)

	room.Session.Store(roomid,roominfo)

	create := &room.Room_create{
		RoomId:roomid,
		RoomMasterId:m.UserId,
		ServerId:conf.ServerId,
		State:0,
		Limit:0,
	}

	return 200,create,""
}



/*
	UserId        string			    `json:"UserId"`
	RoomId 		  string  				`json:"RoomId"`        //房间创建者
	ServerId      int                   `json:"ServerId"`      //服务器对应id
	Sit           int  					`json:"Sit"`           //第一个加入在right 之后再left
*/
func room_join(m *msg.Roomresp) (int ,*room.Room_enter, string){
	r,ok := room.Session.Load(m.RoomId)
	if !ok {
		return 404,nil, fmt.Sprintf("session没有找到roomid:%v",m.RoomId)
	}

	roominfo := r.(room.Room_manager)
	sit ,err:=roominfo.Join(m.UserId,m.RoomId)
	if err != nil  {
		return 500,nil,err.Error()
	}

	join := &room.Room_enter{
		UserId:m.UserId,
		RoomId:m.RoomId,
		ServerId:conf.ServerId,
		Sit:sit,

	}
	return 200,join,""

}



/*
type Room_quit struct {
	UserId        string			    `json:"UserId"`        //退出的用户id
	RoomId 		  string  				`json:"RoomId"`        //房间id
	ServerId      int                   `json:"ServerId"`      //服务器对应id
}
*/

//code 200  301:房间解散,
func room_quit(m *msg.Roomresp) (int ,*room.Room_quit, string) {
	r,ok := room.Session.Load(m.RoomId)
	if !ok {
		return 404,nil, fmt.Sprintf("session没有找到roomid:%v",m.RoomId)
	}
	roominfo := r.(room.Room_manager)


	//删除用户
	code := 200
	limit,err:=roominfo.Quit(m.UserId)
	if err != nil {
		log.Error("bug:用户居然删除不了")
		return 500,&room.Room_quit{},""  //基本上不会发生
	}

	//最后一个用户也退出房间，解散
	if limit == 0 {
		room.Session.Delete(m.RoomId)
		code = 301
	}


	quit := &room.Room_quit{
		UserId:m.UserId,
		RoomId:m.RoomId,
		ServerId:conf.ServerId,
	}

	return code,quit,""

}



/*
type Room_quick_enter struct {
	UserId        string			    `json:"UserId"`
	RoomId 		  string  				`json:"RoomId"`        //房间创建者
	ServerId      int                   `json:"ServerId"`      //服务器对应id
	Sit           int  					`json:"Sit"`
}
*/
func room_quick_enter(m *msg.Roomresp) (int ,*room.Room_quick_enter, string) {
	var roomid string
	var sit int
	var ok bool

	//找到第一个limit
	fn :=func(key, value interface{}) bool {
		roomid = key.(string)
		roominfo := value.(room.Room_manager)

		sit,ok =roominfo.Room_quick_enter()
		if ok {
			return false
		}
		return true
	}

	room.Session.Range(fn)

	quit_enter := &room.Room_quick_enter{
		UserId:m.UserId,
		RoomId:roomid,
		ServerId:conf.ServerId,
		Sit:sit,
	}
	return 200,quit_enter,""
}


