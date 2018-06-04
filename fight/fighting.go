package fight

import (
	"time"
	"errors"
	"server/room"
	"server/conf"
	"server/db"
	"runtime"
	"sync"
	"encoding/json"
	"server/notify"
)

var H = Worker{
	WorkerPool:   make(chan chan Frame, runtime.NumCPU()),
	FrameChannel: make(chan Frame),
}

var FightMap sync.Map
//新建战斗
func NewFighting(c *room.CreateFight) (fight fighting, err error) {
	var num int
	num, err = db.ExistFightByRoomId(c.RoomId)
	if err != nil {
		return
	}
	if num > 0 {
		err = errors.New("该房间已经存在")
		return
	}
	if _, ok := conf.AttributeMap[c.RedRoleId]; !ok {
		err = errors.New("未找到该人物信息")
		return
	}
	if _, ok := conf.AttributeMap[c.BlueRoleId]; !ok {
		err = errors.New("未找到该人物信息")
		return
	}
	fight.Seed = time.Now().UnixNano()
	fight.RoomId = c.RoomId
	fight.MapId = c.MapId
	fight.Red = &roles{
		Uid:       c.RedUid,
		Attribute: conf.AttributeMap[c.RedRoleId],
	}
	fight.Blue = &roles{
		Uid:       c.BlueUid,
		Attribute: conf.AttributeMap[c.BlueRoleId],
	}
	FightMap.Store(c.RoomId, Frame{RoomId: c.RoomId, LastSendTime: time.Now().UnixNano() / 1e6})
	fight.StartTime = time.Now().Unix()
	if err = db.InsertFight(fight); err != nil {
		return
	}
	return

}

//结束战斗
func Close(roomId string, winUid string) error {
	if roomId == "" || winUid == "" {
		return errors.New("参数不完整")
	}
	var result fighting
	err := db.GetFightByRoomId(roomId, &result)
	if err != nil {
		return errors.New("该房间不存在")
	}
	endTime := time.Now().Unix()
	if endTime-result.StartTime < 30 {
		return errors.New("不能再30秒内结束战斗")
	}
	err = db.UpdateFightByRoomId(roomId, winUid, endTime)
	FightMap.Delete(roomId) //删除该房间的信息
	if err != nil {
		return err
	}
	return nil
}

func Run() {
	for {
		H.WorkerPool <- H.FrameChannel
		select {
		case frame := <-H.FrameChannel:
			val, ok := FightMap.Load(frame.RoomId)
			if !ok {
				//房间号丢失,该战斗结束
				notify.SendToRoom("房间号丢失,该战斗结束", frame.RoomId)
				break
			}
			frameData, success := val.(Frame)
			if !success {
				//获取战斗失败,该战斗结束
				FightMap.Delete(frame.RoomId)
				notify.SendToRoom("获取战斗失败,该战斗结束", frame.RoomId)
				break
			}
			frameData.Data = append(frameData.Data, frame.Info)
			if time.Now().UnixNano()/1e6-frameData.LastSendTime >= Interval {
				//转发
				qsort(frameData.Data)
				FightMap.Store(frame.RoomId, Frame{RoomId: frame.RoomId, LastSendTime: time.Now().UnixNano() / 1e6})
				data, _ := json.Marshal(frameData)
				notify.SendToRoom(data, frame.RoomId)
				//TODO 保存每次的信息
			} else {
				//不转发
				FightMap.Store(frame.RoomId, frameData)
			}
			break
		}
	}
}

func qsort(data []FrameData) {
	if len(data) <= 1 {
		return
	}
	key := data[0].Index
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i].Index > key {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}
