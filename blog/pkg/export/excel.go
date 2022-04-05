package export

import "lianxi/blog/pkg/setting"

const EXT = ".xlsx"

//GetExcelFullUrl获取Excel文件的完整访问路径
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

//GetExcelPath获取Excel文件的相对保存路径
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath获取Excel文件的完整保存路径
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
