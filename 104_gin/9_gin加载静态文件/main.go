package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	//加载静态文件可以在url中访问
	r.Static("/aa", "./assets")
	//加载文件下的html文件，可以通过html渲染器返回
	r.LoadHTMLGlob("assets/*")
	r.GET("/index", Go)
	r.Run()
}
func Go(ctx *gin.Context) {
	ctx.HTML(200, "hello.html", nil)
}
