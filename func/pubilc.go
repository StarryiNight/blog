package _func

import (
	"github.com/gin-gonic/gin"
)

type Post struct{//结构体 文章
	Id int
	Title string//标题
	Content string//内容
	Author string//作者
}

//发表文章
func Public (c *gin.Context){
	var post Post
	post.Author=c.PostForm("Author")
	post.Content=c.PostForm("Content")
	post.Title=c.PostForm("Title")
	//读取文章的标题、内容和作者后在数据库创建文章
	post.Create()
}


//在数据库里面新建一个post
func(post *Post) Create()(err error){
	//插入post的sql语句
	sql1 := "insert into posts (title, content, author) values (?, ?, ?)"
	stmt,err:= Db.Prepare(sql1)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res,err:=stmt.Exec(post.Title,post.Content,post.Author)
	if err != nil {
		panic(err)
	}

	//标记文章的ID
	postId,_:=res.LastInsertId()
	post.Id=int(postId)
	return
}


//从数据库删除文章
func Delete(id int)(err error){
	_,err= Db.Exec("delete from posts where id =?",id)
	return
}


