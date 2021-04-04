# 

# 简单博客系统

post:

- /register 注册 检测用户名是否已经存在
-  /login    账户登录 cookie保存60s
- /pubilc   发表文章
       

get:

- /scan       通过文章ID浏览文章

其中pubilc和scan需要登录,若未登录会跳转到登录界面

数据库

post:ID,title,content,author

user:username,password
