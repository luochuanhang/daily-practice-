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
	//设置一个key，过期时间为0，意思是永不过期
	err := client.Set("key", "value", 0).Err()
	//检测有没有设置成功
	if err != nil {
		panic(err)
	}
	//根据key查询缓存，通过Result函数返回两个值
	//第一个代表key，第二个是查询错误信息
	val, err := client.Get("test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
