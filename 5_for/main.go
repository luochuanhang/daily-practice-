package main

import (
	"fmt"
	"time"
)

func main() {
	//最基础的方式，单个for循环
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//经典的初始/条件/后续for循环
	for a := 0; a < 5; a++ {
		fmt.Printf("%d", a)
	}
	//不带条件的循环将一直重复，直到
	//遇到break和return为止
	for {
		fmt.Println("等待5秒")
		time.Sleep(time.Second * 5)
		break
	}
	//可以使用cintine直接进入下一次循环
	for a := 0; a < 10; a++ {
		if a == 5 {
			continue
		}
		fmt.Print(a)
	}

}
