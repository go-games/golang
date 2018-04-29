package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"database/sql"
)


type mysql struct {
	ipaddr string
	db *sql.DB

}

type dbsession struct {
	count uint32
	session map[string]mysql  //改成安全的map
	a uint32
}


func NewDBsession() *dbsession {
	return &dbsession{
		a:100,
		count:5,
		session:make(map[string]mysql),
	}
}


//Mold 取模 区间在100-999之间 超过999 将a 重置到100
//这里没有加锁 使用了原子操作
func (this *dbsession) Mold() uint32{
	if this.a < 999 {
		atomic.AddUint32(&this.a,1)
	}else {
		atomic.StoreUint32(&this.a,100)
	}

	return this.a%this.count

}


func (this *dbsession) MysqlAdd(dbmark,dbaddr string) error {
	db, err :=createMysqlDB()
	if err != nil {
		return err
	}
	this.session[dbmark]= mysql{db:db,ipaddr:dbaddr}
	return nil

}



//// MysqlUpdate 先判断，存在 新建实例替换以前
//func (this *dbsession) MysqlUpdate(dbmark,dbaddr string) {
//	if _, ok := this.session[dbmark]; ok {
//
//	}
//
//}


func (this *dbsession) MysqlDel(dbmark string) {
	if _, ok := this.session[dbmark]; ok {
		delete(this.session, dbmark)
	}
}



func createMysqlDB() (*sql.DB,error) {
	var err error
	db, err := sql.Open("mysql", "root:123456@tcp(10.211.55.4:3306)/game?charset=utf8")
	if err != nil {
		return nil,err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	err = db.Ping()
	if err != nil {
		return nil,err
	}

	return db,nil
}


func main() {
	c:=NewDBsession()
	for {
		v :=c.Mold()
		fmt.Println(v)
		time.Sleep(time.Microsecond*10)
	}
}
