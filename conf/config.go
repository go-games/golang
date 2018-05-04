package conf

import (
	"sync"
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"fmt"
	"reflect"
	"strconv"
	"bytes"
	"golang/model"
	"time"
)




type Server struct {
	Version     string
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
	Mysql       *MysqlInfo
}

//root:123456@tcp(10.211.55.4:3306)/game
type MysqlInfo struct {
	DBname string
	DBaddr string
	DBport string
	DBuser string
	DBpasswd string
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



//第一次运行的时候
func (s *singleton) Load() (err error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	data, err := ioutil.ReadFile("conf/ser.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &s.Server)
	if err != nil {
		log.Fatal("%v", err)
	}

	fmt.Println("ssssssssssssssssss",s.Server)
	return
}




//拼接mysql 字符串
//root:123456@tcp(10.211.55.4:3306)/game
func (s *singleton) MysqlSplice() string {
	var buf bytes.Buffer
	buf.WriteString(s.Server.Mysql.DBuser)
	buf.WriteString(":")
	buf.WriteString(s.Server.Mysql.DBpasswd)
	buf.WriteString("@tcp(")
	buf.WriteString(s.Server.Mysql.DBaddr)
	buf.WriteString(":")
	buf.WriteString(s.Server.Mysql.DBport)
	buf.WriteString(")/")
	buf.WriteString(s.Server.Mysql.DBname)
	buf.WriteString("?charset=utf8")

	return buf.String()
}


//更新失败不能服务不能挂
func (s *singleton) Update(key,value string) error {
	defer func() {
		if err := recover(); err != nil {
			log.Error("--recover-- ", err)
		}
	}()

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


func (this *singleton) Getconf() {
	a:=model.Config{}
//	ist,_:=a.GetConf(0)
//	fmt.Println(ist)
//	a :=make(map[string]interface{})
//	c := make(chan bool)
	go func() {
		timer := time.NewTimer(time.Second * 5)
		for {

			timer.Reset(time.Second * 5)
			select {

			case <-timer.C:
				a.GetVersion()
				if a.Value != this.Server.Version {
					abb:=recursive(0)
					fmt.Println(abb)
					this.SaveConfig("conf/ser.json",abb)
					//this.Server.Version =a.Value
					this.Load()
				}else {
					fmt.Println("版本相同，不用更新",a.Value,this.Server.Version)
				}




			}
		}
	}()


}



//递归实现
func recursive(parent int) map[string]interface{}{
	conf:=model.Config{}
	confList,_:=conf.GetConf(parent)
	a :=make(map[string]interface{})
	for _,aa := range confList {

		if aa.Key == "MaxConnNum" || aa.Key == "ConsolePort"{
			cc ,_ := strconv.Atoi(aa.Value)
			a[aa.Key] = cc
			continue
		}

		a[aa.Key] = aa.Value


		if aa.Struct {
			bb:=recursive(aa.Id)
			a[aa.Key] =bb
		}
	}

	return a
}

//写入json文件
//配置文件保存到本地
func (this *singleton)SaveConfig(filename string,b map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			//log.Error()
		}
	}()

	saveData, _ := json.MarshalIndent(b,"","\t")

	err := ioutil.WriteFile(filename, saveData, 0755)
	log.Debug("保存本地配置文件:",filename)
	if err != nil {
		log.Debug("保存文件失败:",err.Error())
	}
}


//IterConfi 迭代mysql获取的配置文件
//func (this *singleton) RecursiveConf() (a map[string]interface{}){
//	conf:=model.Config{}
//	confList,_:=conf.GetConf()
//
//	a = make(map[string]interface{})
//	for _,aa := range confList {
//		if aa.Parent == 0 {
//			a[aa.Key] = aa.Value
//		}else {
//			key,_ := conf.GetId(aa.Id)
//			a[key] = aa.Value
//		}
//
//	}
//	//
//	fmt.Println(a)
//
//	return
//}




func (this *singleton) JudgeVersion() {
	a := model.Config{}
	a.GetVersion()
	if a.Key == "Version" {
		if a.Value != this.Server.Version {
			//现将配置文件弄成json,存入文件，在说
			this.Server.Version = a.Value
			log.Debug("更新配置文件，", this.Server.Version)



		}else {
			log.Debug("版本号相同，不需要更新配置文件")
		}
	}

}


