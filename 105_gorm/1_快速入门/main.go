package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID         int64
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (User) TableName() string {
	return "users"
}
func main() {
	//配置MySQL连接参数
	username := "root"      //账号
	password := "123456"    //密码
	host := "172.18.110.44" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "db01"        //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，error" + err.Error())
	}
	u := User{
		Username:   "luochan",
		Password:   "123456",
		CreateTime: time.Now().Unix(),
	}
	if err := db.Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	us := User{}
	result := db.Where("username=?", "luochan").First(&us)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		return
	}
	fmt.Println(us.Username, us.Password)
	db.Model(&User{}).Where("username=?", "luochan").Update("password", "654321")
	db.Where("username=?", "luochan").Delete(&User{})
}
