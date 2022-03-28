package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Emp struct {
	Name string `gorm:"column:ENAME"`
}

func (e Emp) TableName() string {
	return "EMP"
}

//mysql dsn格式
//涉及参数:
//username   数据库账号
//password   数据库密码
//host       数据库连接地址，可以是Ip或者域名
//port       数据库端口
//Dbname     数据库名
//username:password@tcp(host:port)/Dbname?charset=utf8&parseTime=True&loc=Local

//填上参数后的例子
//username = root
//password = 123456
//host     = localhost
//port     = 3306
//Dbname   = tizi365
//后面K/V键值对参数含义为：
//  charset=utf8 客户端字符集为utf8
//  parseTime=true 支持把数据库datetime和date类型转换为golang的time.Time类型
//  loc=Local 使用系统本地时区
//root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local

//gorm 设置mysql连接超时参数
//开发的时候经常需要设置数据库连接超时参数，gorm是通过dsn的timeout参数配置
//例如，设置10秒后连接超时，timeout=10s
//root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local&timeout=10s

//设置读写超时时间
// readTimeout - 读超时时间，0代表不限制
// writeTimeout - 写超时时间，0代表不限制
//root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s
func main() {
	db := Db()
	s := "KING"
	e := Emp{}
	//开发的时候需要打开调试日志，这样gorm会打印出执行的每一条sql语句。
	db.Debug().Where("ENAME=?", s).First(&e)
	fmt.Println(e)
}
func Db() *gorm.DB {
	//username = root
	//password = 123456
	//host     = localhost
	//port     = 3306
	//Dbname   = tizi365
	username := "root"
	password := "123456"
	host := "172.18.100.186"
	port := 3306
	Dbname := "db01"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接错误")
	}
	return Db
}
