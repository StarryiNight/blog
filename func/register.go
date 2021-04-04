package _func

import (

	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
func Register(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")

	IsExist(username,password,c)

	if !UserSignup(username,password,c) {
		c.JSON(500,gin.H{
			"status":http.StatusInternalServerError,
			"message":" 数据库Insert出错",
		})
	}else{
		c.JSON(200,gin.H{
			"status":http.StatusOK,
			"message":"注册成功",
		})
	}
}

func IsExist(username, password string, c *gin.Context) bool {
	stmt,err:= Db.Query("select * from user;")
	if err != nil {
		c.JSON(403,"连接数据库失败")
		return false
	}
	defer stmt.Close()
	for stmt.Next() {
		var name,passwd string
		err:=stmt.Scan(&name,&passwd)
		if err != nil {
			c.JSON(403,"查询数据库数据失败")
			return false
		}
		if username == name {
			c.JSON(http.StatusForbidden,gin.H{
				"error":"该用户已经存在!",
			})
			return false
		}
	}
	return true
}

func UserSignup(username string,password string,c *gin.Context )bool{

	_,err:= Db.Exec(
		"INSERT into user(username,password) values(?,?)",username,password)
	if err != nil {
		return  false
	}else{
		return true
	}

}
