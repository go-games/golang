package model


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

func Init(mysqlconn string) {

	var err error
	db, err = sql.Open("mysql", mysqlconn)
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


func Check(err interface{}) {
	if err != nil {
		fmt.Println(err)
	}
}


func getRst(db *sql.DB, sql string) (eles []map[string]interface{}) {
	rst, err := db.Query(sql)
	Check(err)
	columns, err := rst.Columns()
	count := len(columns)

	values := make([]string, count)
	ptr := make([]interface{}, count)
	for rst.Next() {
		for i := 0; i < count; i++ {
			ptr[i] = &values[i]
		}
		rst.Scan(ptr...)
		entry := make(map[string]interface{}, 15)
		for i, col := range columns {
			val := values[i]
			//println(val)
			entry[col] = val
		}
		eles = append(eles, entry)

	}
	return eles
}








