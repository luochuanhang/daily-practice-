package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	//根据redis配置初始化一个客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "172.18.107.135:6379", //redis地址
		Password: "",                    //redis密码，没有则为空
		DB:       0,                     //默认数据库，默认是0
	})
	//订阅channel1这个channel
	sub := client.Subscribe("channell")
	i, err := sub.Receive()
	if err != nil {
		//panic(err)
	}
	switch i.(type) {
	case *redis.Subscription:
	//订阅成功
	case *redis.Message:
		///处理收到消息
		//类型转换
		m := i.(redis.Message)
		//打印收到的消息
		fmt.Println(m)
	case *redis.Pong:
		//收到pong消息
	default:
	}

}
