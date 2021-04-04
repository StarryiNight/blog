package main

import (
	"cookie/func"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



type User struct {
	//用户 :账户名,密码
	Username string`form:"username" json:"username" binding:"required"`
	Password string`form:"password" json:"password" bdinding:"required"`

}


func main()  {

	_func.DatabaseInit()
	//路由
	r:=gin.Default()
	r.POST("/login", _func.Login)                       //登录
	r.POST("/register", _func.Register)                 //注册
	r.POST("/pubilc", _func.LoginCheck(), _func.Public) //发表文章
	r.GET("/scan", _func.LoginCheck(), _func.Scan)      //浏览文章
	r.Run()

}
