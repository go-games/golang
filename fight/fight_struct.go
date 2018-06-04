package fight

import (
	"server/conf"
)

const Interval = 50

//战斗信息
type fighting struct {
	_Id       int    `bson:"_id"`
	Seed      int64  `bson:"seed"`
	RoomId    string `bson:"roomid"`
	MapId     int    `bson:"mapid"`
	WinUID    string `bson:"winuid"`
	Red       *roles `bson:"red"`
	Blue      *roles `bson:"blue"`
	StartTime int64  `bson:"starttime"`
	EndTime   int64  `bson:"endtime"`
}

//用户基础信息
type roles struct {
	Uid       int            `json:"uid"`
	Attribute conf.Attribute `json:"attribute"`
}

type Frame struct {
	RoomId       string       `json:"room_id"`
	Data         []FrameData `json:"data"`
	Info         FrameData   `json:"-"`
	LastSendTime int64       `json:"-"`
}

type FrameData struct {
	Index int64       `json:"index"`
	Order interface{} `json:"order"`
}

type Worker struct {
	WorkerPool   chan chan Frame
	FrameChannel chan Frame
}


