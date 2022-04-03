package logging

import (
	"fmt"
	"time"

	"lianxi/blog/pkg/setting"
)

// 获取日志文件的保存路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// 获取log文件名称
func getLogFileName() string {
	//根据时间格式每天创建一个log日志
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}
