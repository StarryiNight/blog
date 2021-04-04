package _func

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//查询是否为登陆状态,如果未登录则跳转到登录界面
func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,err:=c.Cookie("username")//查询cookie检查是否登录
		if err != nil {

			c.JSON(http.StatusMovedPermanently,"请先登录!")
			c.Redirect(http.StatusMovedPermanently,"/login")//跳转到登录页面
			c.Abort()
		}
	}
}
