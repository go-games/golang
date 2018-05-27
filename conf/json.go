package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}
var MongoCfg struct{
	Username string
	Passwd string
	Domain string
	Port int
	Database string
}
var AttributeMap map[int]Attribute

//角色基础信息
type Attribute struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	PH     int    `json:"ph"`
	Energy int    `json:"energy"`
}
func init() {
	var mData []byte
	data, err := ioutil.ReadFile("conf/server.json")
	mData,err = ioutil.ReadFile("conf/mongod.json")
	aData, err = ioutil.ReadFile("conf/attribute.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	var attributes []Attribute
	err = json.Unmarshal(data, &Server)
	err = json.Unmarshal(mData,&MongoCfg)
	err = json.Unmarshal(aData, &attributes)
	if err != nil {
		log.Fatal("%v", err)
	}
	AttributeMap = make(map[int]Attribute, len(attributes))
	for _, v := range attributes {
		AttributeMap[v.Id] = v
	}
}
