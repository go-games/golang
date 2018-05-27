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
func init() {
	var mData []byte
	data, err := ioutil.ReadFile("conf/server.json")
	mData,err = ioutil.ReadFile("conf/mongod.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	err = json.Unmarshal(mData,&MongoCfg)
	if err != nil {
		log.Fatal("%v", err)
	}
}
