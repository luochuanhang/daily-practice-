package main

import "github.com/go-redis/redis"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.18.107.135:6379", //redis地址
		Password: "",                    //redis密码，没有则为空
		DB:       0,                     //默认数据库，默认是0
	})
	client.Publish("channell", "message")
}
