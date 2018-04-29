package model

import (
	"fmt"
	"strconv"
)

type Config struct {
	Id int
	Key string
	Value string
	Parent int
	Struct bool
}



func (this *Config) TableName() string {
	return TableName("config")
}



func (this *Config) GetVersion() error {
	rows, err := db.Query("SELECT * FROM config where id = 1")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&this.Id,&this.Key,&this.Value,&this.Parent,&this.Struct)
		if err != nil {
			return err
		}
	}
	return nil
}


func (this *Config) GetId(id int) (string,error) {

	fmt.Println("SELECT * FROM config where id ="+strconv.Itoa(id))
	rows, err := db.Query("SELECT * FROM config where id = ?",id)
	if err != nil {
		return "",err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(this.Id,&this.Key,&this.Value,&this.Parent,&this.Struct)
		if err != nil {
			return "",err
		}
	}
	return this.Key,nil
}


func (this *Config) GetConf(parent int) ([]Config,error){
	fmt.Println("SELECT * FROM config where id > 1 and parent = "+strconv.Itoa(parent))
	rows, err := db.Query("SELECT * FROM config where  parent = ?",parent)
	if err != nil {
		return nil,err
	}

	defer rows.Close()
	//checkErr(err)
    var list []Config
	for rows.Next() {
		err = rows.Scan(&this.Id,&this.Key,&this.Value,&this.Parent,&this.Struct)
		if err != nil {
			return nil,err
		}
		//fmt.Println(err)
		//checkErr(err)
		list =append(list,*this)
		//fmt.Println(this)
		//fmt.Println()
		//fmt.Println(department)
		//fmt.Println(created)
	}

	return list,nil
}



