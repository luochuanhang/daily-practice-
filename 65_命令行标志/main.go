package main

import (
	"flag"
	"fmt"
)

//命令行标志是命令行程序指定选项的常用方式。
//例如在wc -l中，这个-l就是一个命令行标志。
func main() {
	//Go 提供了一个 flag 包，支持基本的命令行标志解析
	//基本的标记声明仅支持字符串、整数和布尔值选项
	wordPtr := flag.String("word", "foo", "a string")

	//使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	//用程序中已有的参数来声明一个标志也是可以的。 注意在标志声明函数中需要使用该参数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
