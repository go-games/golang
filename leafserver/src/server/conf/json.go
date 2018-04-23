package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
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

func init() {
	//此处配置文件可以随意选择
	data, err := ioutil.ReadFile("/Users/yang-pc/Go/src/github.com/name5566/golang/leafserver/src/server/conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
		log.Error("lErr", err)

	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Error("lErr2", err)

		log.Fatal("%v", err)
	}
	log.Error("log:data", string(data), Server.ConsolePort)
}
