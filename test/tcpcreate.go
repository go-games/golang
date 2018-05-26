package main

import (
	"encoding/binary"
	"net"
)


/*
	RoomId   string `json:"RoomId"`
	ServerId string `json:"ServerId"`
	UserId string `json:"UserId"`
*/
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体


	/*
		Name    string
		RoomId   string `json:"RoomId"`
		ServerId string `json:"ServerId"`
		UserId string `json:"UserId"`
		ProtoId  string `json:"ProtoId"`  //协议ID 根据id来区分
	*/
	data := []byte(`{
		"Roomresp": {
			"RoomId": "2000",
			"ServerId": "1000",
			"Userid"  : "101",
			"protoId" : 1000
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)
}
