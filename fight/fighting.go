package fight

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"
	"errors"
	"server/room"
	"server/conf"
)

var fightMap sync.Map

//新建战斗
func NewFighting(c *room.CreateFight) ([]byte, error) {
	if _, ok := conf.AttributeMap[c.RedRoleId]; !ok {
		return nil, errors.New("未找到该人物信息")
	}
	if _, ok := conf.AttributeMap[c.BlueRoleId]; !ok {
		return nil, errors.New("未找到该人物信息")
	}
	var fight fighting
	seed := time.Now().Unix()
	fight.Seed = seed
	fight.RoomId = c.RoomId
	fight.MapId = c.MapId
	fight.Keyframe = &Keyframe{Seed: seed, Index: getRandNum(seed, 0), NextIndex: getRandNum(seed, 1)}
	fight.Red = &roles{
		Uid:        c.RedUid,
		BasicState: newBasicState(conf.AttributeMap[c.RedRoleId].PH, conf.AttributeMap[c.RedRoleId].Energy),
		Attribute:  conf.AttributeMap[c.RedRoleId],
	}
	fight.Blue = &roles{
		Uid:        c.BlueUid,
		BasicState: newBasicState(conf.AttributeMap[c.BlueRoleId].PH, conf.AttributeMap[c.BlueRoleId].Energy),
		Attribute:  conf.AttributeMap[c.BlueUid],
	}
	data, err := json.Marshal(fight)
	if err != nil {
		return nil, err
	}
	fightMap.Store(c.RoomId, fight)
	return data, nil

}


//转发战斗数据
func (c *Keyframe) ForwardPlayerData() (bool, error) {
	fight, ok := fightMap.Load(c.RoomId)
	if !ok {
		return false, errors.New("未找到该房间信息")
	}
	data, ok := fight.(fighting)
	if !ok {
		return false, errors.New("获取房间详情失败")
	}
	if c.Index != data.Keyframe.NextIndex {
		if c.Index != getRandNum(data.Keyframe.Seed, c.Frequency) {
			return false, errors.New("当前帧索引不一致")
		}
	}
	data.Keyframe.Frequency = data.Keyframe.Frequency + 1
	data.Keyframe.Index = c.Index
	data.Keyframe.NextIndex = getRandNum(data.Keyframe.Seed, c.Frequency+1)
	if c.Sit == 1 {
		data.Blue.BasicState.SetEnergy(c.Energy)
		data.Blue.BasicState.SetPH(c.PH)
	} else {
		data.Red.BasicState.SetEnergy(c.Energy)
		data.Red.BasicState.SetPH(c.PH)
	}
	fightMap.Store(c.RoomId, data)
	return true, nil
}

//结束战斗
func (c *Keyframe) Close() {
	fightMap.Delete(c.RoomId)
}

func getRandNum(seed int64, num int) (code int) {
	rand.Seed(seed)
	i := 0
	for {
		code = rand.Intn(8888888888)
		if i == num {
			break
		}
		i++
	}
	return
}
