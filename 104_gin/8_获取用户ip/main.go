package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/ip", func(ctx *gin.Context) {
		s := ctx.ClientIP()
		fmt.Println(s)
	})
	r.Run()
}
