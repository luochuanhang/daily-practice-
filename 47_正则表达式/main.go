package main

import (
	"fmt"
	"regexp"
)
//测试一个字符串是否符合一个表达式
func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
}
