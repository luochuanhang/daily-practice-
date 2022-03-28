package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Student struct {
	Name  string   `form:"name"`
	Age   string   `form:"age"`
	Hobby []string `form:"hobby"`
}
type Student2 struct {
	Name  string   `form:"name"`
	Age   string   `form:"age"`
	Hobby []string `form:"hobby"`
}

func main() {
	r := gin.New()
	r.POST("/student", user)
	r.Run()
}
func user(ctx *gin.Context) {
	stu := Student{}
	stu2 := Student2{}
	err := ctx.ShouldBindBodyWith(&stu, binding.JSON)
	if err != nil {
		fmt.Println("1111")
	}
	err = ctx.ShouldBindBodyWith(&stu2, binding.JSON)
	if err != nil {
		fmt.Println("222")
	}
	fmt.Println(stu)
	fmt.Println(stu2)
}
