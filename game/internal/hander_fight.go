package internal

import (
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/utils"
	"server/room"
	"server/fight"
	"server/notify"
)

func handleFight(args []interface{}) {
	m := args[0].(*msg.FightResp)
	a := args[1].(gate.Agent)
	switch m.ProtoId {
	case 2000:
		//开始战斗
		result := StartFight(m)
		if result.Code != 200 {
			a.WriteMsg(&result)
		} else {
			notify.SendToRoom(result.Result, m.RoomId)
		}
	case 2001:
		//提交战斗中玩家指令
		if m.RoomId != "" {
			frame := <-fight.H.WorkerPool
			frame <- fight.Frame{RoomId: m.RoomId, Info: fight.FrameData{Index: m.Index, Order: m.Data}}
		}else {
			a.WriteMsg(utils.Resp_result(404, "","房间号不存在"))
		}
	case 2002:
		//结算战斗
		if err := fight.Close(m.RoomId, m.WinUserId); err != nil {
			a.WriteMsg(utils.Resp_result(404, "", err.Error()))
		} else {
			notify.SendToRoom(m, m.RoomId)
		}
	default:
		log.Error("room协议id不匹配", m.ProtoId)
		a.WriteMsg(utils.ErrGame("protoid not found", "", 404))
	}
}

func StartFight(m *msg.FightResp) (r utils.Result) {
	//开始战斗
	var roomInfo room.Room
	roomInfo = room.NewRoom()
	info, ok := roomInfo.Get(m.RoomId)
	if !ok {
		r.Code = 404
		r.Error = "未找到该房间"
		return
	}
	fight, err := fight.NewFighting(info)
	if err != nil {
		r.Code = 500
		r.Error = err.Error()
		return
	}
	r.Code = 200
	r.Result = fight
	return
}
