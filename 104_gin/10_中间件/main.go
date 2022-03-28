package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//中间件，在请求中间起到拦截作用的处理器函数
//中间件也是一个handler
//每一次请求都会执行中间件
func Zhong1(ctx *gin.Context) {
	fmt.Println("这是第一个")
}
func Zhong2(ctx *gin.Context) {
	s := ctx.ClientIP()
	if s == "127.0.0.1" {
		fmt.Println("这个ip需要等待3秒")
		time.Sleep(3 * time.Second)
	}
	fmt.Println("这是第二个")
}
func hello(ctx *gin.Context) {

	ctx.String(200, "ooo")
}
func main() {
	r := gin.New()
	r.Use(Zhong1, Zhong2)
	r.GET("/hello", hello)
	r.Run()
}
