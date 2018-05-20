package room


import (
	"sync"
	"github.com/gorilla/websocket"
)


//Room/C_Create
/*
    {
     "Error": "",
     "Code": 200,
     "Result": {
      "RoomMasterId": "100001 ",//房间创建者
      "RoomId": "12202120",
      "ServerId": 1000111, //服务器对应id
      "State": 0 // 0 创建成功，1人数已满 2 等待回收或者失效
      "Limit ": "0 " //人数
     }
    }
*/


//创建房间
type Room_create struct {
	RoomId 		  string  				`json:"RoomId"`        //房间创建者
	RoomMasterId  string  				`json:"roomMasterId"`  // 0 创建成功，1人数已满 2 等待回收或者失效
	ServerId      int                   `json:"ServerId"`      //服务器对应id
	State         int
	Limit         int32                `json:"Limit"`         //人数
}




//Room/C_Enter
/*

 "Error": "",
 "Code": 200,
 "Result": {
  "UserId": "100001 ",//进入房间玩家id
  "RoomId": "12202120",
  "ServerId": 1000111, //服务器对应id
  "Sit": 0 //位置 0为左边 1为右边
 }
}
*/

//进入房间
type Room_enter struct {
	UserId        string			    `json:"UserId"`
	RoomId 		  string  				`json:"RoomId"`        //房间创建者
	ServerId      int                   `json:"ServerId"`      //服务器对应id
	Sit           int  					`json:"Sit"`
}




//Room/C_Quit
/*
{
 "Error": "",
 "Code": 200,
 "Result": {
  "UserId": "100001 ",//进入房间玩家id
  "RoomId": "12202120",
  "ServerId": 1000111, //服务器对应id
 }
}
*/

//退出房间
type Room_quit struct {
	UserId        string			    `json:"UserId"`        //退出的用户id
	RoomId 		  string  				`json:"RoomId"`        //房间id
	ServerId      int                   `json:"ServerId"`      //服务器对应id
}



//Room/C_QuickEnter

/*
{
 "Error": "",
 "Code": 200,
 "Result": {
  "UserId": "100001 ",//进入房间玩家id
  "RoomId": "12202120",
  "ServerId": 1000111, //服务器对应id
  "Sit": 0 //位置 0为左边 1为右边
 }
}
*/

//快速进入匹配房间
type Room_quick_enter struct {
	UserId        string			    `json:"UserId"`
	RoomId 		  string  				`json:"RoomId"`        //房间创建者
	ServerId      int                   `json:"ServerId"`      //服务器对应id
	Sit           int  					`json:"Sit"`
}


var Session sync.Map

//type room_session struct {
//	sync.Map   //安全的map
//}


type room_info struct {
	RoomId 		  string  				`json:"RoomId"`        //房间num
	RoomMasterId  string  				`json:"roomMasterId"`  //房间创建者
	ServerId      int                   `json:"ServerId"`      //服务器对应id
	State         int                   `json:"State"`                       // 0 创建成功，1人数已满 2 等待回收或者失效
	Limit         int32                   `json:"Limit"`         //人数
	RoomUsers     sync.Map              `json:"-"`
	//UserId        string			    `json:"UserId"`
}


type room_user_info struct {
	UserId    string
	UserName  string
	Sit       int   //位置 0为左边 1为右边
	conn      *websocket.Conn   //保持的连接
}


