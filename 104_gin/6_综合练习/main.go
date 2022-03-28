package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/hello/:password", zonghe)
	r.Run()

}
func zonghe(ctx *gin.Context) {

	//获取查询参数
	s := ctx.Query("name")
	//获取url参数
	p := ctx.Param("password")
	//默认的查询参数
	l := ctx.DefaultQuery("age", "12")
	//获取表单参数
	f := ctx.PostForm("file")
	fmt.Println(s, p, l, f)

}
