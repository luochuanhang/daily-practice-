package main

import (
	"errors"
	"fmt"
	tools "lianxi/105_gorm/4_tools"

	"gorm.io/gorm"
)

type Food struct {
	Id int
	//常用标签 columb 列名
	//指定主键 primarykey
	//- 忽略字段
	Name       string
	Price      float64
	TypeId     int
	CreateTime int64 `gorm:"column:createtime"`
}

func (f Food) TableName() string {
	return "food"
}
func main() {
	db := tools.GetDB()
	f := Food{}
	//take取   查询一条记录
	//等价于：SELECT * FROM `foods`   LIMIT 1
	db.Debug().Take(&f)
	fmt.Println(f)
	//first 第一根据主键排序查询第一个
	//等价于SELECT * FROM `food` WHERE `food`.`id` = 1 ORDER BY `food`.`id` LIMIT 1
	db.Debug().First(&f)
	fmt.Println(f)
	//last最后 根据主键倒叙查询第一个
	db.Debug().Last(&f)
	fmt.Println(f)
	//find查 查询多条记录
	//因为返回多个数据所以定义切片
	var fs []Food
	db.Debug().Find(&fs)
	fmt.Println(fs)
	//pluck拉,摘 返回一列记录
	var s []string
	db.Model(&Food{}).Debug().Pluck("name", &s)
	fmt.Println(s)
	//当gorm的First、Last、Take 方法找不到记录时会返回 ErrRecordNotFound 错误。
	err := db.Take(&f).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("查询不到数据")
	} else if err != nil {
		//如果err不等于record not found错误，又不等于nil，那说明sql执行失败了。
		fmt.Println("查询失败", err)
	}
	ff := Food{}
	//where 条件
	db.Where("id=?", 2).Debug().Take(&f)
	fmt.Println(ff)
	// in
	db.Where("id in (?)", []int{3, 1}).Debug().Last(&f)
	fmt.Println(ff)
	//在？和?之间
	db.Where("id > ? and id < ?", 1, 3).Debug().Take(&f)
	fmt.Println(ff)
	//like语句
	db.Where("name like ?", "%w%").Debug().Take(&f)
	fmt.Println(ff)

	//select指定返回的条件
	db.Select("name,price").Where("id=?", 2).Debug().Take(&f)
	fmt.Println(ff)
	//Model函数，用于指定绑定的模型，这里生成了一个Food{}变量。
	//目的是从模型变量里面提取表名，Pluck函数我们没有直接传递绑定表名的结构体变量，gorm库不知道表名是什么，所以这里需要指定表名
	//Pluck函数，主要用于查询一列值
	to := []float64{}
	db.Model(&Food{}).Select("count(*)as total").Debug().Pluck("total", &to)
	fmt.Println(to[0])

	//order 设置排序语句，order by子句
	db.Where("price >= ?", 40).Order("price desc").Find(&f)
	fmt.Println(ff)
	//设置limit和Offset子句，分页的时候常用语句。
	fff := []Food{}
	db.Order("price desc").Limit(3).Offset(0).Debug().Find(&fff)
	fmt.Println(fff)
	//Count函数，直接返回查询匹配的行数。
	var total int64 = 0
	//等价于: SELECT count(*) FROM `foods`
	//这里也需要通过model设置模型，让gorm可以提取模型对应的表名
	db.Model(Food{}).Count(&total)
	fmt.Println(total)

	//设置group by子句
	db.Model(Food{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Debug().Scan(&s)
	//Group函数必须搭配Select函数一起使用

	//直接执行sql语句
	//gorm通过db.Raw设置sql语句，通过Scan执行查询。
	sql := "SELECT type, count(*) as  total FROM `food` where price> ? GROUP BY type HAVING (total > 0)"
	db.Raw(sql).Scan(&s)
}
