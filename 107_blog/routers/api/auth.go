package api

import (
	"lianxi/107_blog/models"
	"lianxi/107_blog/pkg/setting/e"
	"lianxi/107_blog/pkg/setting/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//授权结构体
type auth struct { //有效 要求 最大50
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

//授权中间件
func GetAuth(c *gin.Context) {
	//获取URL中的username和password
	username := c.Query("username")
	password := c.Query("password")
	//初始化参数验证
	valid := validation.Validation{}
	//初始化授权结构体
	a := auth{Username: username, Password: password}
	//根据参数校验验证参数，tag标签定义有效的范围
	ok, _ := valid.Valid(&a)
	//创建map存放数据
	data := make(map[string]interface{})
	//错误码默认有错
	code := e.INVALID_PARAMS
	//如果参数验证通过
	if ok {
		//检查这个数据库有没有账号密码
		isExist := models.CheckAuth(username, password)
		//如果有
		if isExist {
			//创建token
			token, err := util.GenerateToken(username, password)

			if err != nil {
				//如果有错设置错误码
				code = e.ERROR_AUTH_TOKEN
			} else {
				//将token数据放入存放数据的map中
				data["token"] = token
				//错误码设置为成功
				code = e.SUCCESS
			}
		} else {
			//如果数据库没有数据设置对应的错误码
			code = e.ERROR_AUTH
		}
	} else {
		//如果参数验证不通过，打印错误
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	//返回前端对应的数据
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
