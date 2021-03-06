package main

import (
	"fmt"
	"log"
	"net/http"

	"lianxi/blog/models"

	"github.com/gin-gonic/gin"

	"lianxi/blog/pkg/gredis"
	"lianxi/blog/pkg/logging"
	"lianxi/blog/pkg/setting"
	"lianxi/blog/pkg/util"
	"lianxi/blog/routers"
)

func init() {
	//参数初始化
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://lianxi/blog
// @license.name MIT
// @license.url https://lianxi/blog/blob/master/LICENSE
func main() {
	//根据字符串设置gin模式
	gin.SetMode(setting.ServerSetting.RunMode)
	//初始化路由
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
