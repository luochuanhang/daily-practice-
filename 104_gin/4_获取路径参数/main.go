package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/hello/:id/:age", getid)
	r.Run()
}

func getid(ctx *gin.Context) {
	s := ctx.Param("id")
	age := ctx.Param("age")
	//获取路劲参数
	fmt.Println(s)
	fmt.Println(age)
}
