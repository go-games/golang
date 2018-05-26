package internal

import (
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/utils"
	"server/room"
	"server/fight"
	"fmt"
)

func handleFight(args []interface{}) {
	m := args[0].(*msg.FightResp)
	a := args[1].(gate.Agent)
	switch m.ProtoId {
	case 2000:
		//开始战斗
		var roomInfo room.Room
		info, ok := roomInfo.Get(200)
		if !ok {
			a.WriteMsg(utils.Resp_result(404, nil, "未找到该房间"))
		}
		figthing, err := fight.NewFighting(info)
		if err != nil {
			a.WriteMsg(utils.Resp_result(500, nil, err.Error()))
		}
		//TODO 发送战斗信息
		fmt.Println(figthing)
	case 2001:
		//提交战斗中玩家指令
	case 2002:
		//结算战斗
	case 2003:
		//开始战斗
	case 2004:
		//退出战斗
	default:
		log.Error("room协议id不匹配", m.ProtoId)
		a.WriteMsg(utils.ErrGame("protoid not found", "", 404))
	}
}
