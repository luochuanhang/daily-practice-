package main

import (
	"fmt"
	"lianxi/107_blog/pkg/setting"
	"lianxi/107_blog/routers"
	"log"
	"syscall"

	"lianxi/107_blog/models"
	"lianxi/107_blog/pkg/logging"

	"github.com/fvbock/endless"
)

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"lianxi/107_blog/pkg/setting"

// 	"lianxi/107_blog/routers"
// )

// func main() {

// 	router := routers.InitRouter()

// 	s := &http.Server{
// 		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
// 		Handler:        router,
// 		ReadTimeout:    setting.ReadTimeout,
// 		WriteTimeout:   setting.WriteTimeout,
// 		MaxHeaderBytes: 1 << 20,
// 	}

// 	go func() {
// 		if err := s.ListenAndServe(); err != nil {
// 			log.Printf("Listen: %s\n", err)
// 		}
// 	}()
// 	//“Signal”表示操作系统信号。通常的底层实现是依赖于操作系统的:
// 	//在Unix上，它是sycall . signal。
// 	quit := make(chan os.Signal, 1)
// 	/*
// 		Notify使包信号将传入信号转发给chan。如果没有提供信号，所有传入的信号将被中继到c。
// 		否则，只有提供的信号才会。
// 	*/
// 	//信号.通知             中断
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit
// 	//打印日志
// 	log.Println("Shutdown Server ...")

// 	/*
// 		context.Background() 返回一个空的Context
// 		我们可以用这个空的Context作为goroutine的root节点（如果把整个goroutine的关系看作树状）
// 		使用context.WithCancel(parent)函数，创建一个可取消的子Context
// 		函数返回值有两个：子Context Cancel取消函数
// 	*/
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	//使用 Shutdown 可以优雅的终止服务，其不会中断活跃连接。
// 	//工作过程是关闭所有的开启的监听器，然后关闭所有闲置的连接，
// 	//最后等待活跃的连接均闲置了才终止服务
// 	if err := s.Shutdown(ctx); err != nil {
// 		//打印日志并退出
// 		log.Fatal("Server Shutdown:", err)
// 	}
// 	//打印日志
// 	log.Println("Server exiting")
// }
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
