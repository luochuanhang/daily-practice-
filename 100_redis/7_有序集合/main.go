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
	//1.ZAdd 添加一个或者多个元素到集合，如果元素已经存在则更新分数
	// 添加一个集合元素到集合中， 这个元素的分数是2.5，元素名是tizi
	// err := client.ZAdd("age2", redis.Z{3, "tt"}).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//2.ZCard
	//返回集合元素个数
	size, err := client.ZCard("age2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	//3.ZCount
	//统计某个分数范围内的元素个数
	// 返回： 1<=分数<=5 的元素个数, 注意："1", "5"两个参数是字符串
	size, err = client.ZCount("age2", "3", "5").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	// 返回： 1<分数<=5 的元素个数
	// 说明：默认第二，第三个参数是大于等于和小于等于的关系。
	// 如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
	size, err = client.ZCount("age2", "(1", "5").Result()
	fmt.Println(size)

	//4.ZIncrBy
	//增加元素的分数
	//给元素5增加2分
	client.ZIncrBy("age2", 2, "5")

	//5.ZRange,ZRevRange
	//返回集合中某个索引范围的元素，根据分数从小到大排序
	// 返回从0到-1位置的集合元素， 元素按分数从小到大排序
	// 0到-1代表则返回全部数据
	vals, err := client.ZRange("age2", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	for _, val := range vals {
		fmt.Println(val)
	}
	//ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。

	//6.ZRangeByScore
	//根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	// 初始化查询条件， Offset和Count用于分页
	op := redis.ZRangeBy{
		Min:    "2",  // 最小分数
		Max:    "10", // 最大分数
		Offset: 0,    // 类似sql的limit, 表示开始偏移量
		Count:  5,    // 一次返回多少数据
	}

	vals, err = client.ZRangeByScore("age2", op).Result()
	if err != nil {
		panic(err)
	}

	for _, val := range vals {
		fmt.Println(val)
	}
	//7.ZRevRangeByScore
	//用法类似ZRangeByScore，区别是元素根据分数从大到小排序。

	//8.ZRangeByScoreWithScores
	//用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
	// 初始化查询条件， Offset和Count用于分页
	op = redis.ZRangeBy{
		Min:    "2",  // 最小分数
		Max:    "10", // 最大分数
		Offset: 0,    // 类似sql的limit, 表示开始偏移量
		Count:  5,    // 一次返回多少数据
	}

	val2, err := client.ZRangeByScoreWithScores("age2", op).Result()
	if err != nil {
		panic(err)
	}

	for _, val := range val2 {
		fmt.Println(val.Member) // 集合元素
		fmt.Println(val.Score)  // 分数
	}

	//8.ZRem
	//删除集合元素
	// 删除集合中的元素tizi
	client.ZRem("age2", "1")
	// 删除集合中的元素tizi和xiaoli
	// 支持一次删除多个元素
	client.ZRem("age2", "3", "5")

	//9.ZRemRangeByRank
	//根据索引范围删除元素
	// 集合元素按分数排序，从最低分到高分，删除第0个元素到第5个元素。
	// 这里相当于删除最低分的几个元素
	client.ZRemRangeByRank("age2", 0, 5)

	// 位置参数写成负数，代表从高分开始删除。
	// 这个例子，删除最高分数的两个元素，-1代表最高分数的位置，-2第二高分，以此类推。
	client.ZRemRangeByRank("age2", -1, -2)

	//10.ZRemRangeByScore
	//根据分数范围删除元素
	// 删除范围： 2<=分数<=5 的元素
	client.ZRemRangeByScore("key", "2", "5")
	// 删除范围： 2<=分数<5 的元素
	client.ZRemRangeByScore("key", "2", "(5")

	//11.ZScore
	//查询元素对应的分数
	// 查询集合元素tizi的分数
	score, _ := client.ZScore("key", "tizi").Result()
	fmt.Println(score)

	//12.ZRank
	//根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
	rk, _ := client.ZRank("key", "tizi").Result()
	fmt.Println(rk)
	//ZRevRank的作用跟ZRank一样，区别是ZRevRank是按分数从大到小排序。
}
