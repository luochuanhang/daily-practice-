package main

import (
	"fmt"
	tools "lianxi/105_gorm/4_tools"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name string
}

func (t Test) TableName() string {
	return "test"
}

//通过AutoMigrate自动建表
func main() {
	db := tools.GetDB()
	db.AutoMigrate(&Test{})

	//检测表是否存在
	//可以通过结构体和表名
	b := db.Migrator().HasTable("USER")
	fmt.Println(b)

	//根据Test结构体建表
	db.Migrator().CreateTable(&Test{})

	//删除表
	//db.Migrator().DropTable(&Test{})

	//删除结构体中的字段
	db.Migrator().DropColumn(&Test{}, "name")
}
