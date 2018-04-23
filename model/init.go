package model


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(10.211.55.4:3306)/game?charset=utf8")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

}

func TableName(name string) string {
	//return ("db.prefix") + name
	return name
}

