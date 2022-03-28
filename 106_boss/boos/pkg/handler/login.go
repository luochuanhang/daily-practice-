package handler

import (
	"fmt"
	"net/http"
	"time"

	"lianxi/106_boss/boos/pkg/model"
	"lianxi/106_boss/boos/pkg/storage"
	"lianxi/106_boss/boos/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginHandler struct {
	base
}

func (l *LoginHandler) Init(router *gin.RouterGroup, engine storage.Storage) {
	l.engine = engine
	l.logger = logrus.WithField("handler", "login")
	router.POST("/login", l.Login)
}

func (l *LoginHandler) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		l.badrequest(ctx)
		return
	}
	db := l.engine.Get().Debug()
	var user model.User
	result := db.Where("username = ? and password = ?",
		req.Username,
		util.Password(req.Password)).Find(&user)
	if result.Error != nil {
		l.logger.WithError(result.Error).Errorf("failed to get account with username %s", req.Username)
		l.badrequest(ctx)
		return
	}
	if result.RowsAffected != 1 {
		l.badrequest(ctx)
		l.logger.Infof("user not found %s", req.Username)
		return
	}
	ctx.SetCookie("boos_session", fmt.Sprintf("%d", user.ID), int(7*time.Hour), "/", "*", false, false)
	if err := storage.Get().Set(fmt.Sprintf("session_%d", user.ID), user.ID, 10*time.Minute).Err(); err != nil {
		l.logger.WithError(err).Errorf("faild to write to redis")
		l.serverError(ctx)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "success login",
	})
}
