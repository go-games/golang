package logic

import (

	"crypto/md5"
	"io"
	"fmt"
	"server/db"
)

func CheckUser(u string,passwd string) (status bool,err error)  {
	h := md5.New()
	io.WriteString(h,passwd)
	mdstr := fmt.Sprintf("%x",h.Sum(nil))
	db.Table("mvm_member_table")
	db.Where("member_id","=",u)
	userInfo := db.Find()
	if userInfo["member_name"] == u && userInfo["member_pass"] == mdstr {
		return true,nil
	}else {
		return  false,fmt.Errorf("用户名或密码错误")
	}

}