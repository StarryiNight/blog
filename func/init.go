package _func

import "database/sql"

var Db *sql.DB

func DatabaseInit(){
	//连接数据库
	var err error
	Db, err = sql.Open("mysql", "root:123456@/user?charset=utf8")
	if err != nil {
		panic(err)
	}
}
