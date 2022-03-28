package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/login", func(c *gin.Context) {
		//第一个参数cookie的名字
		//第二个参数cookie的值
		//第三个参数过期时间，单位秒
		//第四个参数cookie所在位置
		//第五个参数访问域名
		//第六个参数是否只能通过https访问
		//第7个参数是否能被js获取
		c.SetCookie("abc", "123", 60, "/", "127.0.0.1", false, true)
		c.String(200, "登录成功")
	})
	r.GET("/home", AuthMiddlWare(), func(c *gin.Context) {
		c.String(200, "欢迎来到home")
	})
	r.Run()
}
func AuthMiddlWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取cookie的值
		cook, err := ctx.Cookie("abc")
		fmt.Println(cook)
		if err == nil {
			if cook == "123" {
				//执行下一个handler
				ctx.Next()
				return
			}
		}
		//还没有cookie不能访问
		ctx.String(http.StatusUnauthorized, "未知")
		//不在执行后续的handler
		ctx.Abort()
	}
}
