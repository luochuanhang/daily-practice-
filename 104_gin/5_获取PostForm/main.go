package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/hello", postfo)
	r.Run()
}
func postfo(ctx *gin.Context) {
	//获取表单参数
	s := ctx.PostForm("name")
	age := ctx.PostForm("age")
	//多选框参数
	hoby := ctx.PostFormArray("hoby")

	fmt.Println(s)
	fmt.Println(age)
	fmt.Println(hoby)
}
