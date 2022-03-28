package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type student struct {
	Name string `form:"name"`
	Age  string `form:"agg"`
}

func main() {
	//注册服务
	e := gin.New()
	//session引擎
	store := cookie.NewStore([]byte("1111"))
	//添加中间件
	g := e.Group("/v1")
	g.Use(sessions.Sessions("mysession", store))
	//加载静态文件
	g.Static("assets", "assets/")
	//html绑定页面
	e.LoadHTMLGlob("assets/*")
	//登录
	g.POST("/login", login)
	//查看
	g.GET("/home", Anth(), home)
	e.Run()

}
func login(ctx *gin.Context) {
	name := ctx.PostForm("name")
	pass := ctx.PostForm("password")
	fmt.Println(name)
	session := sessions.Default(ctx)
	ctx.SetCookie(name, name, 600, "/", "127.0.0.1", false, true)
	session.Set(name, pass)
	session.Save()
	ctx.HTML(200, "hello.html", 200)
	ctx.String(200, "登录成功")
	stu := student{}
	ctx.ShouldBind(&stu)
	fmt.Println(stu)
}
func Anth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")
		coo, err := ctx.Cookie(name)
		if err == nil {
			fmt.Println(name)
			fmt.Println(coo)
			if coo == name {
				ctx.Next()
				return
			}

		}
		ctx.Abort()
		ctx.String(200, "未登录")
	}
}
func home(ctx *gin.Context) {
	name := ctx.Query("name")
	ctx.String(200, "欢迎来到", name, "的home")
}
