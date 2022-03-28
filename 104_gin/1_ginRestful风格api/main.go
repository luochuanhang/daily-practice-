package main

import "github.com/gin-gonic/gin"

func main() {
	//注册路由
	r := gin.New()
	//restful可以根据相同的路径不同的请求方法处理不同的功能
	//绑定路径和handler
	r.GET("/book", func(ctx *gin.Context) {
		ctx.String(200, "这是GET")
	})
	r.POST("/book", func(ctx *gin.Context) {
		ctx.String(200, "这是POST")
	})
	r.PUT("/book", func(ctx *gin.Context) {
		ctx.String(200, "这是PUT")
	})
	r.DELETE("/book", func(ctx *gin.Context) {
		ctx.String(200, "这是DELETE")
	})
	//默认8080端口
	r.Run()
}
