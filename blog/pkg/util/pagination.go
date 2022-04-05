package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"lianxi/blog/pkg/setting"
)

// GetPage获取页面参数
func GetPage(c *gin.Context) int {
	result := 0
	//获取URI参数转换为int
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
