package _func

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

)

//登录
func Login(c *gin.Context) {

	username:=c.PostForm("username")
	password:=c.PostForm("password")
	//如果查询数据库账户名密码正确返回true 否则false
	if UserSignin(username,password){
		c.SetCookie(//设置cookie "username"
			"username", username,
			60, //有效时间60s
			"/",
			"localhost",
			false,
			true,
		)
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"message":"登陆成功",
		})
	}else {//登陆失败
		c.JSON(403,gin.H{
			"status":http.StatusForbidden,
			"message":"登陆失败,密码错误或账户名不存在",
		})
	}
}

//从数据库查询数据比对账户密码
func UserSignin(username string,password string)bool  {
	stmt,err:= Db.Query("select * from user;") //查询的sql语句
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	for stmt.Next(){//遍历比对账户密码
		var name string
		var word string
		err:=stmt.Scan(&name,&word)
		if err != nil {
			fmt.Println(err)
		}

		if username==name&&password==word{
			return true
		}
	}
	return false
}
