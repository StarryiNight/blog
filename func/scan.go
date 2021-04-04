package _func

import (

	"github.com/gin-gonic/gin"
	"strconv"
)

func Scan(c *gin.Context){

	id,_:=strconv.Atoi(c.PostForm("ID"))
	post,_:= GetPost(id)
	c.JSON(200,gin.H{
		"Title":post.Title,
		"Author":post.Author,
		"Content":post.Content,
	})
}

func GetPost(id int)(post Post,err error){
	post= Post{}
	err= Db.QueryRow("select id,title,content,author from posts where id =?",id).
		Scan(&post.Id,&post.Title,&post.Content,&post.Author)
	return
}