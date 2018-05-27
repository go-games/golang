package fight

import (
	"time"
	"errors"
	"server/room"
	"server/conf"
	"server/db"
)

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
	if err != nil {
		return err
	}
	return nil
}
