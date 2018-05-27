package db

import (
	"fmt"
	"github.com/name5566/leaf/db/mongodb"
	"server/conf"
	"gopkg.in/mgo.v2"
	"encoding/json"
)
var C *mongodb.DialContext
var s *mongodb.Session
func init()  {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",conf.MongoCfg.Username,conf.MongoCfg.Passwd,conf.MongoCfg.Domain,conf.MongoCfg.Port,conf.MongoCfg.Database);
     dbContext,err := mongodb.Dial(url,200)
    if err != nil  {
    	panic(err)
	}
    C = dbContext
    s = C.Ref()
}

func Insert(data interface{})error{
	defer C.UnRef(s)
	err :=s.DB(conf.MongoCfg.Database).C("player").Insert(data)
	if err != nil && err != mgo.ErrNotFound {
		return err
	 }
	return nil
	}
func FindOnce(condition interface{},result *interface{})error{
	defer C.UnRef(s)
	r_bytes,_ := json.Marshal(condition)
	err := s.DB(conf.MongoCfg.Database).C("player").Find(condition).One(result)
	if err != nil {
		return fmt.Errorf("mongodDB错误:%s",err)
	}
	if err == mgo.ErrNotFound {
		return fmt.Errorf("指定的条件:%s,数据未找到匹配项",string(r_bytes))
	}
	return nil
}

