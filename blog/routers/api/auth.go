package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"lianxi/blog/pkg/app"
	"lianxi/blog/pkg/e"
	"lianxi/blog/pkg/util"
	"lianxi/blog/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	//参数校验
	valid := validation.Validation{}
	//获取表单账号密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	//创建验证实例
	a := auth{Username: username, Password: password}
	//验证一个结构体。可以在结构体中添加tag标签设置参数要求
	ok, _ := valid.Valid(&a)

	if !ok {
		//MarkErrors记录错误日志
		app.MarkErrors(valid.Errors)
		//返回json数据
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	//创建服务器验证实例
	authService := auth_service.Auth{Username: username, Password: password}
	//检查账号密码是否存在
	isExist, err := authService.Check()
	if err != nil {
		//返回json数据
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		//返回json数据
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	//生成token令牌
	token, err := util.GenerateToken(username, password)
	if err != nil {
		//返回json数据
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	//返回json数据
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
