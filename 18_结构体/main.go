package main

import "fmt"

//结构体是带类型的字段集合
type person struct {
	name string
	age  int
}

//newPerson使用给定的名字构造一个新的person结构体
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 28
	//可以返回指定局部变量的指针，因为局部变量将在函数
	//的作用域中继续存在
	return &p
}

func main() {
	//使用这个语法可以创建新的结构体元素
	fmt.Println(person{"luo", 15})
	//可以在初始化一个结构体元素是指定字段名字
	//省略的字段将被初始化为零值
	fmt.Println(person{name: "chuan"})
	fmt.Println(person{age: 20})
	//&前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 18})
	//在构造函数中封装创建新的结构
	fmt.Println(newPerson("Jon"))
	//使用.来访问结构体字段。
	s := person{name: "sean", age: 50}
	fmt.Println(s.name)
	//也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

}
