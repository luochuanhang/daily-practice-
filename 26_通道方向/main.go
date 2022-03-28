package main

import "fmt"

//ping 函数定义了一个只允许发送数据的通道
func ping(ch chan<- string, str string) {
	ch <- str
}

//pong函数允许通道pings来接收数据，
//另外一个通道来发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pin := make(chan string, 1)
	pon := make(chan string, 1)
	ping(pin, "nihao")
	pong(pin, pon)
	fmt.Println(<-pon)
}
