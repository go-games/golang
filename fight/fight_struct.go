package fight

//战斗信息
type fighting struct {
	Seed     int64     `json:"seed"`
	RoomId   int       `json:"room_id"`
	MapId    int       `json:"map_id"`
	Red      *roles    `json:"red"`
	Blue     *roles    `json:"blue"`
	Keyframe *Keyframe `json:"keyframe"`
}

//角色基础信息
type attribute struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	PH     int    `json:"ph"`
	Energy int    `json:"energy"`
}

//用户基础信息
type roles struct {
	Uid        int         `json:"uid"`
	BasicState *basicState `json:"basic_state"`
	Attribute  attribute   `json:"attribute"`
}

//帧信息
type Keyframe struct {
	Seed        int64       `json:"seed"`        //随机数种子
	Frequency   int64       `json:"frequency"`   //调用次数
	Index       int         `json:"index"`       //当前帧索引
	NextIndex   int         `json:"next_index"`  //下一帧索引
	Instruction interface{} `json:"instruction"` //客户端指令
}

