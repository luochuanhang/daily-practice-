package main

import (
	"fmt"
)

//常规的通道发送和接收数据是阻塞的，我们可以通过
//default子句的select来实现非阻塞的发送，接收，
//甚至是非阻塞的多路select
func main() {
	messages := make(chan string)
	signal := make(chan bool)
	//如果在 messages 中存在，然后 select 将这个
	//值带入 <-messages case中。如果不是，
	//就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	//可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
