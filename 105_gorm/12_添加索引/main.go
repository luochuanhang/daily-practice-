package main

import (
	tools "lianxi/105_gorm/4_tools"

	"gorm.io/gorm"
)

//两个字段使用同一个索引名，Migration将创建复合索引
type User struct {
	gorm.Model
	Name   string `gorm:"index:idx_member"`
	Number string `gorm:"index:idx_member"`
}

func (u User) TableName() string {
	return "user"
}

func main() {
	db := tools.GetDB()
	//创建数据库
	db.Migrator().CreateTable(&User{})
	//添加索引
	db.Migrator().CreateIndex(&User{}, "name")
	db.Migrator().CreateIndex(&User{}, "idx_name")
	// 为 Name 字段删除索引
	db.Migrator().DropIndex(&User{}, "name")
	db.Migrator().DropIndex(&User{}, "idx_name")
	// 检查索引是否存在
	db.Migrator().HasIndex(&User{}, "name")
	db.Migrator().HasIndex(&User{}, "idx_name")
}
