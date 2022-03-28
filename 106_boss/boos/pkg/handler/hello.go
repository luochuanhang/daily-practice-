package handler

import (
	"lianxi/106_boss/boos/pkg/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Hello struct {
	base
}

func (h *Hello) Init(router *gin.RouterGroup, engine storage.Storage) {
	h.engine = engine
	h.logger = logrus.WithField("handler", "hello")
	hellogroup := router.Group("/hello")
	hellogroup.GET("/", h.Hello)
}

func (h *Hello) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "hello!",
	})
}
