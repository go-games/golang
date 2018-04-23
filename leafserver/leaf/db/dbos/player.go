package dbos

import (
	"github.com/name5566/leaf/db/mysql"
)

//ExistPlayer 判断用户是否存在
func ExistPlayer(username string) (exist bool, err error) {
	var pw string
	if err := mysql.GetPlayerStmt.QueryRow(username).Scan(&pw); err != nil {
		return false, err
	}
	if pw != "" {
		return true, nil
	}
	return false, nil
}

//InsertPlayer  增加玩家
func InsertPlayer(username string, pw string, salt string, unixtime int64) (err error) {
	if _, err := mysql.AddPlayerStmt.Exec(username, pw, salt, unixtime); err != nil {
		return err
	}
	return nil
}
