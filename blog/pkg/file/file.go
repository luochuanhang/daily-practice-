package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize 回去文件的大小
func GetSize(f multipart.File) (int, error) {
	//读取数据直到出现err和EOF
	content, err := ioutil.ReadAll(f)
	//根据读取数据的长度判断获得文件的大小
	return len(content), err
}

// GetExt返回文件的后缀
func GetExt(fileName string) string {
	//返回最后一个最后一个斜杠，最后一个点后面的字段
	return path.Ext(fileName)
}

// checkknotexist检查文件是否存在
func CheckNotExist(src string) bool {
	//查看文件信息
	_, err := os.Stat(src)
	//返回文件是否不存在
	return os.IsNotExist(err)
}

// CheckPermission检查文件是否有权限
func CheckPermission(src string) bool {
	//查看文件状态
	_, err := os.Stat(src)
	//返回错误是否是无权访问的错误
	return os.IsPermission(err)
}

// IsNotExistMkDir创建一个不存在的目录
func IsNotExistMkDir(src string) error {
	//检查文件是否不存在				是
	if notExist := CheckNotExist(src); notExist {
		//创建目录
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir创建目录
func MkDir(src string) error {
	//MkdirAll创建一个名为path的目录，以及任何必要的父目录，并返回nil，否则返回一个错误。
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open根据特定的模式进行文件归档
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	//文件不存在就根据给定模式创建，并返回file用于io
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen 尝试打开文件
func MustOpen(fileName, filePath string) (*os.File, error) {
	//返回与当前目录对应的根路径名
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	//根路径+文件地址
	src := dir + "/" + filePath
	//检查这个目录是否无权限访问
	perm := CheckPermission(src)

	if perm {
		//不能访问
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	//创建一个不存在的目录
	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	//打开文件，不存在就创建
	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
