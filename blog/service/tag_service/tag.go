package tag_service

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"

	"lianxi/blog/models"
	"lianxi/blog/pkg/export"
	"lianxi/blog/pkg/file"
	"lianxi/blog/pkg/gredis"
	"lianxi/blog/pkg/logging"
	"lianxi/blog/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

//检查是否有相同的标签
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

//检查标签是否存在
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

//添加标签
func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

//编辑标签
func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

//删除标签
func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

//标签总数
func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

//获取tag
func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)
	//创建tag实例
	cache := cache_service.Tag{
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	//获取key
	key := cache.GetTagsKey()
	//检查redis是否存在
	if gredis.Exists(key) {
		//获取key中的数据
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			//将data反序列化为[]tag
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	//获得标签
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	//将数据存入redis数据库
	gredis.Set(key, tags, 3600)
	return tags, nil
}

//导出标签
func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}
	//创建一个新文件
	xlsFile := xlsx.NewFile()
	//AddSheet用提供的名称将一个新表添加到文件中。表名的最小长度为1个字符。
	//如果工作表名称长度较短，则抛出错误。
	sheet, err := xlsFile.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	//向工作表添加新行
	row := sheet.AddRow()
	/*
		Cell是一个高级结构，用于在xlsx中为用户提供对Cell内容的访问。行。
	*/
	var cell *xlsx.Cell
	for _, title := range titles {
		//将字符串的字符添加title
		cell = row.AddCell()
		cell.Value = title
	}
	//遍历标签
	for _, v := range tags {
		values := []string{
			//将int转换为字符串
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}
		//向工作表添加新行
		row = sheet.AddRow()
		for _, value := range values {
			//向这一个行添加数据
			cell = row.AddCell()
			cell.Value = value
		}
	}
	//将时间戳转换为字符串
	time := strconv.Itoa(int(time.Now().Unix()))
	//文件的名字tags+时间戳+文件后缀
	filename := "tags-" + time + export.EXT
	//获取Excel的保存路径
	dirFullPath := export.GetExcelFullPath()
	//如果目录不存在就创建一个
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}
	//将文件保存到提供的路径下的xlsx文件中。
	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

//导入标签
func (t *Tag) Import(r io.Reader) error {
	//OpenReader把一个io返回一个填充的XLSX文件。
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	//GetRows根据给定的工作表名称(区分大小写)返回工作表中的所有行。
	rows := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			data = append(data, row...)
			//添加标签
			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}

//没有删除有名字有状态信息
func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
