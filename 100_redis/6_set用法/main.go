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
	//1.SAdd
	//添加集合元素
	err := client.SAdd("age1", 100).Err()
	if err != nil {
		panic(err)
	}
	client.SAdd("age1", 100, 200, 300, 555, 666, 777)

	//2.SCard
	//获取集合元素个数
	size, err := client.SCard("age1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	//3.SIsMember
	//判断元素是否在集合中
	// 检测100是否包含在集合中
	ok, _ := client.SIsMember("age1", 100).Result()
	if ok {
		fmt.Println("集合包含指定元素")
	}

	//4.SMembers
	//获取集合中所有的元素
	es, _ := client.SMembers("age1").Result()
	// 返回的es是string数组
	fmt.Println(es)

	//5.SRem
	//删除集合元素
	// 删除集合中的元素100
	client.SRem("age1", 100)

	// 删除集合中的元素tizi和2019
	client.SRem("age1", 200, 300)

	//SPop,SPopN 随机返回集合中的元素，并且删除返回的元素
	// 随机返回集合中的一个元素，并且删除这个元素
	val, _ := client.SPop("age1").Result()
	fmt.Println(val)

	// 随机返回集合中的5个元素，并且删除这些元素
	vals, _ := client.SPopN("age1", 5).Result()
	fmt.Println(vals)

}
