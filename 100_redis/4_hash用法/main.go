package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.18.107.135:6379",
		Password: "",
		DB:       0,
	})
	//Hset 根据key和field字段设置，field字段的值
	err := client.HSet("user_1", "username", "tizi365").Err()
	if err != nil {
		panic(err)
	}

	//HGet 根据key和field字段，查询field字段的值
	username, err := client.HGet("user_1", "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(username)

	//HGetAll
	//根据key查询所有字段值
	// 一次性返回key=user_1的所有hash字段和值
	data, err := client.HGetAll("user_1").Result()
	if err != nil {
		panic(err)
	}
	// data是一个map类型，这里使用使用循环迭代输出
	for field, val := range data {
		fmt.Println(field, val)
	}

	//HIncrBy 根据key和field字段，累加字段的数值
	//累加count字段的值，一次性累加2， user_1为hash key
	count, err := client.HIncrBy("user_1", "count", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	//HKeys 根据key返回所有字段名
	// keys是一个string数组
	keys, err := client.HKeys("user_1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	//HLen 根据key，查询hash的字段数量
	size, err := client.HLen("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	//HMGet 根据key和多个字段名，批量查询多个hash字段值
	//HMGet支持多个field字段名，意思是一次返回多个字段值
	vals, err := client.HMGet("key", "field1", "field2").Result()
	if err != nil {
		panic(err)
	}
	// vals是一个数组
	fmt.Println(vals)

	//HMSet  根据key和多个字段名和字段值，批量设置hash字段值
	// 初始化hash数据的多个字段值
	// 初始化hash数据的多个字段值
	data1 := make(map[string]interface{})
	data1["id"] = 1
	data1["username"] = "tizi"

	// 一次性保存多个hash字段值
	err = client.HMSet("key", data1).Err()
	if err != nil {
		panic(err)
	}

	//HSetNX
	//如果fieId字段不存在，则设置hash字段值
	err = client.HSetNX("key", "id", 100).Err()
	if err != nil {
		panic(err)
	}
	//HDel
	//根据key和字段名，删除hash字段，支持批量删除hash字段
	//删除一个字段
	client.HDel("key", "id")
	//删除多个字段
	client.HDel("key", "id", "username")

	//HExists
	//检测hash字段名是否存在
	// 检测id字段是否存在
	err = client.HExists("key", "id").Err()
	if err != nil {
		panic(err)
	}

}
