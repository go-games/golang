package main

import (
	"encoding/binary"
	"flag"
	"net"
)


/*
type Room_resp struct {
	RoomId   string `json:"RoomId"`
	ServerId string `json:"ServerId"`
	UserId string `json:"UserId"`
}


*/

func main() {

	var version int
	var addr string
	flag.IntVar(&version, "count", 10000, "print version")
	flag.StringVar(&addr, "addr", "127.0.0.1:3563", "服务器端地址")
	flag.Parse()

	data := []byte(`{
		"Room_create": {
			"RoomId": "2001",
			"ServerId": "123456"
			"UserId": "56"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)


	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}


	conn.Write(m)

}
