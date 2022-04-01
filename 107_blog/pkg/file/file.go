package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//获取文件大小
func GetSize(f multipart.File) (int, error) {
	/*
		ReadAll从r读取数据，直到出现错误或EOF，并返回所读取的数据。
		一个成功的调用返回err = = nil，而不是err = = EOF。
		因为ReadAll被定义为从src读取到EOF，所以它不会将从read读取
		的EOF视为要报告的错误。
	*/
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

//获取文件后缀
func GetExt(fileName string) string {
	/*
		Ext返回path使用的文件扩展名。扩展名是path的最后一个斜杠分隔元素中
		从最后一个点开始的后缀;如果没有点，则为空。
	*/
	return path.Ext(fileName)
}

//检查文件是否存在
func CheckNotExist(src string) bool {
	/*
		Stat返回一个描述命名文件的filelinfo。如果有错误，它的类型将是*PathError。
	*/
	_, err := os.Stat(src)
	/*
		IsNotExist返回一个布尔值，指示该错误是否已知，以报告某个文件或目录不存在。
		它由ErrNotExist和一些系统调用错误所满足。这个函数出现在errors.ls之前。
		它只支持操作系统包返回的错误。新代码应该使用errors.lslerr。fs.ErrNotExist)。
	*/
	return os.IsNotExist(err)
}

//检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	//判断文件是否是无权限的错
	return os.IsPermission(err)
}

//如果文件不存在就新建文件夹
func IsNotExistMkDir(src string) error {
	//判断文件是否存在
	if notExist := CheckNotExist(src); notExist {
		//不存在就创建文件
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

//新建文件夹
func MkDir(src string) error {
	/*
		MkdirAll创建一个名为path的目录，以及任何必要的父目录，并返回nil，
		否则返回一个错误。权限位perm (umask之前)用于MkdirAll创建的所有目录。
		如果path已经是一个目录，则MkdirAll不执行任何操作并返回nil。
	*/
	/*
		定义的文件模式位是FileMode中最重要的位。最低有效位是标准Unix rwxrwxrwx权限。
		这些位的值应该被认为是公共API的一部分，可以在有线协议或磁盘表示中使用:
		它们不能被更改，尽管可能会添加新的位。
	*/
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	/*
		OpenFile是通用的开放调用;大多数用户将使用Open或Create代替。
		它用指定的标志(O_RDONLY等)打开指定的文件。如果该文件不存在，
		并且传递了O_CREATE标志，则使用模式perm创建它(在umask之前)。
		如果成功，返回的File上的方法可以用于I/O。如果有错误，它的类型将是*PathError。
	*/
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
