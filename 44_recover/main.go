package main

import "fmt"

//Go 通过使用 recover 内置函数，可以从panic中_恢复recover_。
// recover 可以阻止 panic 中止程序，并让它继续执行。

/*
当其中一个客户端连接出现严重错误，服务器不希望崩溃。
相反，服务器希望关闭该连接并继续为其他的客户端提供
服务。 实际上，这就是Go的net/http包默认对于HTTP
服务器的处理。
*/
func mayPanic() {
	panic("a problem")
}
func main() {
	//必须在 defer 函数中调用 recover。 当跳出引发 panic 的函数时，defer 会被激活， 其中的 recover 会捕获 panic。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	//这行代码不会执行，因为 mayPanic 函数会调用 panic。 main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束。
	fmt.Println("After mayPanic()")
}
