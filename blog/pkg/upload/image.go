package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"lianxi/blog/pkg/file"
	"lianxi/blog/pkg/logging"
	"lianxi/blog/pkg/setting"
	"lianxi/blog/pkg/util"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName获取图像名称
func GetImageName(name string) string {
	//获取路径最后一个/字段最后一个.后面的字段
	ext := path.Ext(name)
	//TrimSuffix返回不包含所提供的后缀字符串的s。
	fileName := strings.TrimSuffix(name, ext)
	//对文件名进行MD5编码
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath获取保存路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath获取完整保存路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt检查图像文件
func CheckImageExt(fileName string) bool {
	//返回文件的后缀
	ext := file.GetExt(fileName)
	//文件的后缀和图片后缀进行比较
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if allowExt == ext {
			return true
		}
	}

	return false
}

// CheckImageSize检查图像大小
func CheckImageSize(f multipart.File) bool {
	//获取文件的大小
	size, err := file.GetSize(f)
	if err != nil {
		//输出日志
		log.Println(err)
		logging.Warn(err)
		return false
	}
	//文件大小是否小于或等于设置大小
	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage检查文件是否存在
func CheckImage(src string) error {
	//Getwd返回与当前目录对应的根路径名
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	//检查目录是否存在
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	//检查是否有权限
	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
