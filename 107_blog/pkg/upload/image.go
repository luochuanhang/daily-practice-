package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"lianxi/107_blog/pkg/file"
	"lianxi/107_blog/pkg/logging"
	"lianxi/107_blog/pkg/setting"
	"lianxi/107_blog/pkg/setting/util"
)

//获取图片完整访问 URL
func GetImageFullUrl(name string) string {

	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

//获取图片名称
func GetImageName(name string) string {
	/*
		Ext返回path使用的文件扩展名。扩展名是path的最后一个斜杠
		分隔元素中从最后一个点开始的后缀;如果没有点，则为空。
	*/
	ext := path.Ext(name)
	/*
		TrimSuffix返回不包含所提供的后缀字符串的s。
		如果s不以后缀结尾，则返回s不变。
	*/
	fileName := strings.TrimSuffix(name, ext)
	//将文件名字转换为md5
	fileName = util.EncodeMD5(fileName)
	//返回md5编码的文件名+文件后缀
	return fileName + ext
}

//获取图片路径
func GetImagePath() string {
	//获取文件的路径的保存路径
	return setting.AppSetting.ImageSavePath
}

//获取图片完整路径
func GetImageFullPath() string {

	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

//检查图片后缀
func CheckImageExt(fileName string) bool {
	//返回文件后缀
	ext := file.GetExt(fileName)
	//遍历图片后缀
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		//检查文件名是否是图片
		if strings.EqualFold(ext, allowExt) {
			return true
		}
	}

	return false
}

//检查图片大小
func CheckImageSize(f multipart.File) bool {
	//返回文件的大小
	size, err := file.GetSize(f)
	//有错写入日志returnfalse
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	//检查文件大小有没有超过设置最大值
	return size <= setting.AppSetting.ImageMaxSize
}

//检查图片
func CheckImage(src string) error {
	/*
		Getwd返回与当前目录对应的根路径名。如果当前目录
		可以通过多个路径到达(由于符号链接)，Getwd可能会
		返回其中任何一个。
	*/
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	//检查文件地址存不存在
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	//检查能不能读取文件
	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
