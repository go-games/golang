package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/log"
	"fmt"
)
var db *sql.DB
var err error
var param struct{
	table string
	where string
	field string
	action func()

}
func init() {
	db,err = sql.Open("mysql","root:@tcp(127.0.0.1:3306)/oto178?charset=utf8")
    checkError(err)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	checkError(err)
}
func main()  {
   Where("member_id='admin'")
   Table("mvm_member_table")
   fmt.Printf("%v",Find())
}
func checkError(err error){
	if err != nil {
		log.Error("db error: %s",err)
		panic(err)
	}
}
func Where(s string){
	param.where = s
}
func Table(t string){
    param.table = t
}
//查询单个数据
func Find() map[string]interface{}  {
	var query string
	if param.where == "" || param.table == "" {
		log.Error("table or where is nill")
		panic("table or where is nill")
	}
	query = "SELECT * FROM "+param.table+" WHERE "+param.where+" LIMIT 1"
	rows,err := db.Query(query)
	defer rows.Close()
	checkError(err)
	return  parseRows(rows)
}
//解析结果
func parseRows( rows *sql.Rows)map[string]interface{}{
	fields,_ := rows.Columns()
	var result = make(map[string]interface{},len(fields))
	var values = make([][]byte,len(fields))
	var scans = make([]interface{},len(fields))
	for i := range scans {
		scans[i] = &values[i]
	}

	for rows.Next(){
		rows.Scan(scans...)
		for k, v := range values{
			result[fields[k]] = v

		}
	}
	return result

}
func Field(f string)  {
	param.field = f
}