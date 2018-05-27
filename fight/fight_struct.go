package fight

import (
	"server/conf"
)

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
