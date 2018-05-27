package notify

import (
	"github.com/fanliao/go-concurrentMap"
	"github.com/name5566/leaf/log"
	"time"
	"fmt"
	"github.com/name5566/leaf/gate"
)
type onlineUser struct {
	Conn gate.Agent
	UserLab string//用户标识
	RoomId string //房间标识号
	State bool//是否已处理
}
var Pool *concurrent.ConcurrentMap = concurrent.NewConcurrentMap()//并发安全的map
type key struct{
	lab string //用户标识
	roomid string //房间标识号 已废弃
	on int //上线时间戳
}
//新的用户进入
func New(lab string,agent2 gate.Agent)error{
	v := new(onlineUser)
	v  = &onlineUser{
		Conn:agent2,
		UserLab:lab,
		State:false,
	}
	k := &key{lab,"",timestamp()}
	value,err := Pool.Put(k,v)
	log.Debug("%v",value)
	if err !=nil{
		return err
	}
	if value != nil {
		Pool.Replace(k,v)
	}
	agent2.SetUserData(lab)
	return nil
}
//重载用户配置
func Reload(lab string,agent2 gate.Agent) error  {
	v := &onlineUser{
		Conn:agent2,
		UserLab:lab,
		State:false,
	}
	k := &key{lab:lab,on:1,roomid:""}
	_,err := Pool.Remove(k)
	if err != nil {
		return err
	}
	_,err = Pool.Put(k,v)
	if err !=nil {
		return err
	}
   return nil
}
//订阅信息，暂时只支持房间 待完善
func Subscribe(lab string,roomid string) error {
	k := &key{lab:lab,on:1,roomid:""}
	v, err := Pool.Get(k)
	if err != nil{
		return err
	}
	if v == nil {
		return fmt.Errorf("无相关用户数据")
	}
	v.(*onlineUser).RoomId = roomid
	k = &key{lab:lab,on:timestamp(),roomid:""}//刷新上线时间
	value,err := Pool.Replace(k,v)
	if err !=nil {
		return err
	}
	log.Debug("%v",value)
	return nil
}
//转发给特定用户的信息
func SendOne(msg interface{},lab string)error{
	k := &key{lab:lab,on:1,roomid:""}
	v, err := Pool.Get(k)
	if err != nil || v == nil {
		return err
	}
	v.(*onlineUser).Conn.WriteMsg(msg)
	return nil
}
//转发给服务器上的所有用户
func SendAny(msg interface{}){
	defer func() {
		if err :=recover();err != nil  {
			log.Error("%s",err)
		}
	}()
	for itr := Pool.Iterator();itr.HasNext();{
		k,v,ok :=itr.Next()
		if ok {
			if k.(*key).roomid == "" {
				v.(*onlineUser).Conn.WriteMsg(msg)
			}
		}
	}
}
//转发给特定房间用户的信息
func SendToRoom(msg interface{},roomid string){
	defer func() {
		if err :=recover(); err != nil {
			log.Error("%s",err)
		}
	}()
	for itr := Pool.Iterator();itr.HasNext();{
		_,v,ok :=itr.Next()
		if ok {
			if v.(*onlineUser).RoomId == roomid {
				log.Debug("找到房间号: %s",roomid)
				v.(*onlineUser).Conn.WriteMsg(msg)
			}
		}
	}

}
//释放过期的用户数据，减少内存使用
func Free(){
    go func() {
		for _ = range time.NewTicker(1200*time.Second).C{
			log.Debug("定时任务启动中....")
			now := int(time.Now().Unix())
			for itr := Pool.Iterator(); itr.HasNext(); {
				k,_,ok := itr.Next()
				if ok {
					if k.(*key).on < now {
						log.Debug("%v 清除成功",k)
						Pool.Remove(k)
					}
				}

			}

		}
	}()
}
//过期时间 暂定40分钟后过期
func timestamp() int {
	return int(time.Now().Unix()+2400)
}
func (u *key) HashBytes() []byte {
	return []byte(u.lab)
}
func (u *key) Equals(v2 interface{}) (equal bool) {
	u2, ok := v2.(*key)
	return ok && u.lab == u2.lab
}
