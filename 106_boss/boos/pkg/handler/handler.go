package handler

import (
	"lianxi/106_boss/boos/pkg/storage"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Init(router *gin.RouterGroup, engine storage.Storage)
}
