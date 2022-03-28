package router

import (
	"github.com/gin-gonic/gin"

	"lianxi/106_boss/boos/pkg/config"
	"lianxi/106_boss/boos/pkg/handler"
	"lianxi/106_boss/boos/pkg/middleware"
	"lianxi/106_boss/boos/pkg/storage"
)

var handlers = []handler.Handler{
	&handler.LoginHandler{},
	&handler.RegistHandler{},
	&handler.Hello{},
	&handler.ResumeHandler{},
}

func Start(engine storage.Storage) error {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(middleware.NewAuthenticate()())
	{
		for _, handler := range handlers {
			handler.Init(v1, engine)
		}
	}
	return router.Run(config.DefaultConfig.Addr)
}
