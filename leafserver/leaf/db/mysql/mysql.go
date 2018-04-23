package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/conf"
)

var DB *sql.DB

func init() {
	var err error
	du := conf.GetString("dbusername", "")
	dp := conf.GetString("dbpassword", "")
	da := conf.GetString("dbipport", "")
	dn := conf.GetString("dbname", "")

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?readTimeout=15s&writeTimeout=15s&timeout=10s&charset=utf8mb4", du, dp, da, dn))
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(2000)
	//TODO 不同功能分表
	AddPlayerStmt, err = DB.Prepare("insert into player(username,password,salt,registertime) values(?,?,?,?)")
	printErr(err)
	GetPlayerStmt, err = DB.Prepare("select id from player where username = ?")
	printErr(err)
}
func printErr(err error) {
	if err != nil {
		panic(err)
	}
}

var AddPlayerStmt *sql.Stmt
var GetPlayerStmt *sql.Stmt
