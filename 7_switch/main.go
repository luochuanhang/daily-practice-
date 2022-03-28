package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	//一个最基本的swtich
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	//在同一个case中，可以使用逗号来分隔多个表达式，
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is a weekday")
	}

	//不带表达式的switch是实现if/else
	//逻辑的另一种方式，这里还展示了case
	//表达式也可以不使用常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}
	//类型开关 type swtich比较类型而非值
	//可以用来发现一个接口值得类型，
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I an a int")
		default:
			fmt.Printf("Don know type %T\n", i)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("nihao")

}
