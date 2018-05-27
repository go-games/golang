package fight

import "server/conf"

//战斗信息
type fighting struct {
	Seed     int64     `json:"seed"`
	RoomId   int       `json:"room_id"`
	MapId    int       `json:"map_id"`
	Red      *roles    `json:"red"`
	Blue     *roles    `json:"blue"`
	Keyframe *Keyframe `json:"keyframe"`
}

//用户基础信息
type roles struct {
	Uid        int            `json:"uid"`
	BasicState *basicState    `json:"basic_state"`
	Attribute  conf.Attribute `json:"attribute"`
}

//帧信息
type Keyframe struct {
	Seed        int64       `json:"seed"`        //随机数种子
	Frequency   int         `json:"frequency"`   //调用次数
	Index       int         `json:"index"`       //当前帧索引
	NextIndex   int         `json:"next_index"`  //下一帧索引
	PH          int         `json:"ph"`          //生命值 eg:+1,-1
	Energy      int         `json:"energy"`      //怒气值 eg:+1,-1
	Instruction interface{} `json:"instruction"` //客户端指令

	RoomId int `json:"-"` //房间号
	Sit    int `json:"-"` //1:blue,2:red
}
