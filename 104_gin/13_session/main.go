package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//创建基于cookie的引擎
	store := cookie.NewStore([]byte("1111"))
	//添加session中间件，第一个session名字，第二个session引擎
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/login", func(ctx *gin.Context) {
		//初始化session对象
		session := sessions.Default(ctx)
		//设置session
		session.Set("hello", "nihao")
		//保存session
		session.Save()
	})
	r.GET("/home", Anth(), func(ctx *gin.Context) {
		ctx.String(200, "登录成功了")
	})
	r.Run()
}
func Anth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//创建session对象
		session := sessions.Default(c)
		//如果session值相等
		if session.Get("hello") == "nihao" {
			//执行后面的handler
			c.Next()
			return
		}
		//后面的handler不执行
		c.String(200, "未登录")
		c.Abort()
	}

}
