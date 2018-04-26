package conf

import (
	"sync"
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"fmt"
	"reflect"
	"strconv"
)




type Server struct {
	LogLevel    string  "loglevel"
	LogPath     string  "logpath"
	WSAddr      string  "wsaddr"
	CertFile    string  "certfile"
	KeyFile     string  "keyfile"
	TCPAddr     string  "tcpaddr"
	MaxConnNum  int     "maxconnnum"
	ConsolePort int     "consoleport"
	ProfilePath string  "profilepath"
	Debug       bool
}




type singleton struct {
	Mu      sync.Mutex
	Server  *Server

}


var instance *singleton
var once sync.Once


func  GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}


//类型转发
func (s *singleton) Tostr2int(value string) int {
	v ,err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err.Error())
	}
	return v
}

//第一次运行的时候
func (s *singleton) Load() (err error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &s.Server)
	if err != nil {
		log.Fatal("%v", err)
	}

	return
}


//更新失败不能服务不能挂
func (s *singleton) Update(key,value string) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	server := reflect.ValueOf(s.Server).Elem()

   //判断类型,
	switch server.FieldByName(key).Kind() {
	case reflect.Bool:
		valueBool, err:=strconv.ParseBool(value)
		if err != nil {
			return err
		}
		server.FieldByName(key).SetBool(valueBool)
	//case reflect.Uint64:
	//case reflect.int64:
		//reflect.ValueOf(value).Int()
	case reflect.Int:
		valueInt,err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		if key == "MaxConnNum" {
			s.Server.MaxConnNum = valueInt
		}

		if key == "ConsolePort" {
			s.Server.ConsolePort = valueInt
		}
	case reflect.String:
		server.FieldByName(key).SetString(value)
	//case reflect.Map:
	default:
		return  fmt.Errorf("invalid type: %v %s", server.Kind())
	}
	fmt.Println(server.FieldByName(key))
	return nil
}
