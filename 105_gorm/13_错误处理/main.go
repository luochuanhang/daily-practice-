package main

import (
	"errors"
	tools "lianxi/105_gorm/4_tools"

	"gorm.io/gorm"
)

type User struct {
}

func (u User) TableName() string {
	return "user"
}

//如果在执行SQL查询的时候，出现错误，GORM 会将错误信息
//保存到 *gorm.DB 的Error字段，
//我们只要检测Error字段就可以知道是否存在错误。
func main() {
	db := tools.GetDB()
	if err := db.Where("name = ?", "tizi365").First(&User{}).Error; err != nil {
		// 错误处理
	}
	//ErrRecordNotFound error
	//当 First、Last、Take 方法找不到记录时，GORM
	//会返回 ErrRecordNotFound 错误。如果发生了多个
	//错误，你可以通过 errors.Is 判断错误是否为
	//ErrRecordNotFound
	// 检查错误是否为 RecordNotFound
	err := db.First(&User{}, 100).Error
	b := errors.Is(err, gorm.ErrRecordNotFound)
	if b {
		//处理
	}
}
