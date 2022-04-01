package logging

import (
	"fmt"
	"lianxi/107_blog/pkg/file"
	"lianxi/107_blog/pkg/setting"
	"os"
	"time"
)

//获取日志文件地址
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

//获取日志文件名字
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		//设置的日志文件地址
		setting.AppSetting.LogSaveName,
		//现在的时间
		time.Now().Format(setting.AppSetting.TimeFormat),
		//设置的日志文件后缀
		setting.AppSetting.LogFileExt,
	)
}

//获取全部的日志地址
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

//打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {
	/*
		Getwd返回与当前目录对应的根路径名。如果当前目录可以通过多个路径到达(由于符号链接)，
		Getwd可能会返回其中任何一个。
	*/
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		/*
			Errorf根据格式说明符格式化，并将字符串作为满足错误的值返回。
			如果格式说明符包含一个带有错误操作数的%w谓词，则返回的错误将实现一个
			返回该操作数的Unwrap方法。包含一个以上%w谓词或为其提供一个没有实现错误
			接口的操作数是无效的。动词%w是%v的同义词。
		*/
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
