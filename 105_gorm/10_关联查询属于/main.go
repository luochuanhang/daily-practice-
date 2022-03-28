package main

import (
	"fmt"
	tools "lianxi/105_gorm/4_tools"

	"gorm.io/gorm"
)

//GORM的关联查询（又叫连表查询）中的属于关系是一对一关联关系的一种，通常用于描述一个Model属于另外一个Model。
/*
存在一个users表和profiles表：
users - 用户表
profiles - 用户个性化信息表
他们之间存在一对一关系，每一个用户都有自己的个性化数据，
那么可以说每一条profiles记录都属于某个用户。
*/
type User struct {
	gorm.Model
	Name string
}
type profile struct {
	gorm.Model
	UserID uint //外键
	// 定义user属性关联users表，默认情况使用 类型名 + ID
	// 组成外键名，在这里UserID属性就是外键
	User User
	Name string
}

//定义外键
type Profile1 struct {
	gorm.Model
	Name      string
	User      User `gorm:"foreignkey:UserRefer"` //使用 UserRefer 作为外键
	UserRefer uint // 外键
}

//关联外键
type User1 struct {
	gorm.Model
	Refer string // 关联外键
	Name  string
}

type Profile struct {
	gorm.Model
	Name      string
	User      User1 `gorm:"references:Refer"` // 使用 Refer 作为关联外键
	UserRefer string
}

func main() {
	profile := Profile{}
	db := tools.GetDB()
	// 查询用户个性数据
	//自动生成sql： SELECT * FROM `profiles` WHERE id = 1 AND `profiles`.`deleted_at` IS NULL LIMIT 1
	db.Where("id = ?", 1).Take(&profile)
	fmt.Println(profile)

	user := User{}
	// 通过Profile关联查询user数据, 查询结果保存到user变量
	db.Model(&profile).Association("User").Find(&user)
	fmt.Println(user)
	// 自动生成sql: SELECT * FROM `users` WHERE `users`.`id` = 1 // 1 就是user的 ID，已经自动关联
}
