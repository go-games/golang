package model


//https://www.cnblogs.com/tsiangleo/p/4483657.html 效率

type User struct {
	Id int
	User string
	Passwd string
}



func (u *User) TableName() string {
	return TableName("user")
}


func (u *User) UserGetByName() error{
	rows, err := db.Query("select id,user,passwd from user where user = ? limit 1", u.User)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Id, &u.User,&u.Passwd)
		if err != nil {
			return err
		}
	}
	//num, err := res.RowsAffected()
	return nil
}



func (u *User) UserAdd() (int64,error) {
	rows, err := db.Prepare("INSERT INTO user(user,passwd) values(?,?)")
	if err != nil {
		return 0,err
	}
	defer rows.Close()

	res,err := rows.Exec(u.User,u.Passwd)
	if err != nil {
		return 0,nil
	}

	id,err := res.LastInsertId()

	return id,nil
}
