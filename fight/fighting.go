package fight

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"math/rand"
	"sync"
	"time"
	"errors"
	"server/room"
)

var AttributeMap map[int]attribute
var once sync.Once

func init() {
	once.Do(readAttribute)
}

func readAttribute() {
	var attributes []attribute
	data, err := ioutil.ReadFile("C:\\Project\\Go\\src\\server\\fight\\attribute.json")
	if err != nil {
		log.Panic(err)
		return
	}
	if err = json.Unmarshal(data, &attributes); err != nil {
		log.Panic(err)
		return
	}
	AttributeMap = make(map[int]attribute, len(attributes))
	for _, v := range attributes {
		AttributeMap[v.Id] = v
	}
	log.Println("角色信息初始化完成")
}


//新建战斗
func NewFighting(c *room.CreateFight) ([]byte, error) {
	if _, ok := AttributeMap[c.RedRoleId]; !ok {
		return nil, errors.New("未找到该人物信息")
	}
	if _, ok := AttributeMap[c.BlueRoleId]; !ok {
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
		BasicState: newBasicState(AttributeMap[c.RedRoleId].PH, AttributeMap[c.RedRoleId].Energy),
		Attribute:  AttributeMap[c.RedRoleId],
	}
	fight.Blue = &roles{
		Uid:        c.BlueUid,
		BasicState: newBasicState(AttributeMap[c.BlueRoleId].PH, AttributeMap[c.BlueRoleId].Energy),
		Attribute:  AttributeMap[c.BlueUid],
	}
	data,err:=json.Marshal(fight)
	if err!=nil{
		return nil,err
	}
	return data, nil

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
