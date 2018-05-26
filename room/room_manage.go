package room

import (
	"github.com/name5566/leaf/log"
	"sync/atomic"
	"fmt"
	"strconv"
	"time"
)
type Room interface {
	Get(roomid int) (*CreateFight,bool)   //获取房间信息
}


type Room_manager interface {
	Create(userid string) string
	Join(userid,roomid string) (int,error)
	Quit(userid string) (int32,error)
	Room_quick_enter() (int,bool)
	Close() bool
}


func NewRoomInfo() *room_info {
	r := &room_info{
		State:0,
		Limit:0,
		//RoomUsers:[2]*room_user_info{},
	}
	return r
}

//编译的时候检查
var _ Room_manager = new(room_info)

func (r *room_info) Create(userid string) string {
	r.RoomMasterId = userid
	roomid := gen_room_id()
	log.Debug("创建房间成功,roomid:",roomid)
	log.Debug("插入mongo room")
	return roomid
}



func (r *room_info) Join(userid,roomid string) (int,error) {
	sit := 0
	switch atomic.LoadInt32(&r.Limit) {
	case 0:
	case 1:
		//查看已经存在的用户他的sit，并设置新加入用户sit为另一个值
		b,ok:=r.RoomUsers.Load(r.RoomMasterId)
		if ok {
			if sit == b.(room_user_info).Sit {
				sit = 1
			}
		}else {
			log.Error("bug:房间主人没有在房间里")
			return -2,fmt.Errorf("bug:房间主人没有在房间里")
		}

	default:
		log.Error(userid,"加入失败,房间人数已经满了",roomid)
		return -1,fmt.Errorf("加入失败,房间人数已经满了")

	}
	r.RoomUsers.Store(userid,&room_user_info{UserId:userid,Sit:sit})
	atomic.AddInt32(&r.Limit,1)

	log.Debug(fmt.Sprint("用户",userid," 加入房间成功",roomid," 发送房间广播消息"))
	log.Debug("更新mongo roominfo")

	return sit,nil
}


func (r *room_info) Quit(userid string) (int32,error) {
	r.RoomUsers.Delete(userid)
	atomic.AddInt32(&r.Limit,-1)
	log.Debug(userid,"退出房间成功",r.RoomId," 房间人数:",r.Limit,"发送房间广播消息")
	limit := atomic.LoadInt32(&r.Limit)
	if limit <= 0 {
		log.Debug("关闭房间,删除房间信息",r.RoomId)
		log.Debug("删除mongo roominfo")
	}else {
		log.Debug("更新mongo roominfo")
		//TODO 更新sit信息  设置最后一个人为房主

		var roomMasterId string
		if r.RoomMasterId == userid {
			fn := func(key, value interface{}) bool {
				roomMasterId = key.(string)
				//value.(room_user_info).Sit = 0
				return false
			}

			//返回false才停止 循环
			r.RoomUsers.Range(fn)
		}

		r.RoomMasterId = roomMasterId

	}

	return limit,nil
}


//返回一个 最先找到的只有一个人的房间号
func (r *room_info) Room_quick_enter() (int,bool){
	if atomic.LoadInt32(&r.Limit) == 1 {
		var sit = 0
		//查看已经存在的用户他的sit，并设置新加入用户sit为另一个值
		b,ok:=r.RoomUsers.Load(r.RoomMasterId)
		if ok {
			if sit == b.(room_user_info).Sit {
				sit = 1
			}
			return sit,true
		}else {
			log.Error("bug:房间主人没有在房间里")
			return -1,false
		}
	}
	return -1,false
}


func (r *room_info) Close() bool {
	log.Debug("关闭房间,删除房间信息",r.RoomId)
	log.Debug("删除mongo roominfo")
	return true
}


//给战斗模式的接口 用来获取room info  （从room_session获取）
func (c *CreateFight) Get(roomid int) (*CreateFight,bool) {
	log.Debug("获取房间信息成功")
	r,ok := Session.Load(strconv.Itoa(roomid))
	if !ok {
		return nil,false
	}
	roominfo := r.(room_info)

	var reduid string
	fn := func(key, value interface{}) bool {
		if roominfo.RoomMasterId != key.(string) {
			reduid = value.(room_user_info).UserId
			return false
		}
		return true
	}

	//返回false才停止 循环
	roominfo.RoomUsers.Range(fn)

	c.RoomId = roomid
	c.MapId = MapID
	c.RedUid,_ = strconv.Atoi(reduid)
	c.RedRoleId = 1
	c.BlueUid,_ = strconv.Atoi(roominfo.RoomMasterId)
	c.BlueRoleId =0

	return c,true
}

func NewRoom() *CreateFight {
	return &CreateFight{}
}



//按照纳秒生产一个房间号
func gen_room_id() string {
	return strconv.FormatInt(time.Now().UnixNano(),10)
}