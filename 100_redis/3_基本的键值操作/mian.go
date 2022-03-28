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
	//Set 设置一个键值
	// 第三个参数代表key的过期时间，0代表不会过期。
	err := client.Set("age", 12, 0).Err()
	if err != nil {
		panic(err)
	}

	//Get 查询key的值
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	s, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	//GetSet 设置一个key值，并返回这个key的旧值
	s2, err := client.GetSet("name", "chuan").Result()
	if err != nil {
		panic(err)
	}
	//设置一个key的值，并返回这个key的旧值
	fmt.Println("name", s2)

	//SetNX 如果key不存在，则设置这个key的值
	err = client.SetNX("age", "11", 0).Err()
	if err != nil {
		panic(err)
	}

	//MGet 批量查询key的值
	//MGet函数可以传入任意个key，一次性返回多个值
	//这里result返回两个值，第一个值是一个数组，第二个值是错误信息
	vals, err := client.MGet("name", "test", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

	//Mset 批量设置key的值
	err = client.MSet("key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		panic(err)
	}

	//Incr,IncrBy 针对一个key的数值进行递增操作
	//Incr函数每次加一
	val, err := client.Incr("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最新值", val)
	//IncrBy函数，可以指定每次递增多少
	val, err = client.IncrBy("age", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最新值", val)
	// IncrByFloat函数，可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
	val1, err := client.IncrByFloat("age", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最新值", val1)

	//Decr,DecrBy 对一个key的数值进行递减操作
	// Decr函数每次减一
	val, err = client.Decr("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最新值", val)

	// DecrBy函数，可以指定每次递减多少
	val, err = client.DecrBy("age", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("最新值", val)

	//Del  删除key操作，支持批量删除
	client.Del("age")
	err2 := client.Del("name", "test").Err()
	if err2 != nil {
		panic(err)
	}
	//设置过期时间
	client.Expire("key", 3)
}
