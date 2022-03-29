package main

import (
	"fmt"
	"lianxi/107_blog/pkg/setting"
	"lianxi/107_blog/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := routers.InitRouter()
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test",
		})
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
