package main

import (
	"github.com/gin-gonic/gin"
)

var student = gin.H{
	"luo":  gin.H{"name": "luochuanghang", "age": "52"},
	"zhan": gin.H{"name": "zhan", "age": "44"},
}

func main() {
	r := gin.New()
	g := r.Group("/v1", gin.BasicAuth(gin.Accounts{
		"luo":  "123",
		"zhan": "123",
	}))
	g.GET("/user", admin)
	r.Run()
}
func admin(ctx *gin.Context) {
	user := ctx.MustGet(gin.AuthUserKey).(string)
	if stu, ok := student[user]; ok {
		ctx.String(200, user, stu)
	} else {
		ctx.String(200, "meiyou")
	}
}
