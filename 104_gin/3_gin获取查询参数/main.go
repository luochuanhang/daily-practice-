package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//注册服务
	r := gin.New()
	//绑定路由和handler
	r.GET("/book", getbook)
	//启动服务
	r.Run()
}

func getbook(ctx *gin.Context) {
	//获取url中查询参数对应的value
	//c.request.url.Query
	s := ctx.Query("name")
	num := ctx.Query("age")
	//获取url查询参数中的字段值，如果没有值设置一个默认值
	a := ctx.DefaultQuery("password", "123")
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(num)
}
