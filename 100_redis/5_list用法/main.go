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
	//LPush 从列表左边插入数据
	client.LPush("name", "luo")
	//LPush支持一次插入任意个数据
	err := client.LPush("name", 1, 2, 3, 4).Err()
	if err != nil {
		panic(err)
	}

	//LPushX
	//跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。

	//RPop
	//从列表的右边删除第一个数据，并返回删除的数据
	val, err := client.RPop("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	//RPush
	//从列表右边插入数据
	client.RPush("name", "data1")
	// 支持一次插入任意个数据
	err = client.RPush("name", 1, 2, 3, 4, 5).Err()
	if err != nil {
		panic(err)
	}

	//LPop 从列表左边删除第一个数据，并返回删除的数据
	val, err = client.LPop("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	//LLen
	//返回列表的大小
	i, err := client.LLen("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	//LRange 返回列表的一个范围内的数据，也可以返回全部数据
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err := client.LRange("name", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

	//LRem  删除列表中的数据
	// 从列表左边开始，删除100， 如果出现重复元素，仅删除1次，也就是删除第一个
	_, err = client.LRem("name", 1, 100).Result()
	if err != nil {
		panic(err)
	}

	// 如果存在多个100，则从列表左边开始删除2个100
	client.LRem("key", 2, 100)

	// 如果存在多个100，则从列表右边开始删除2个100
	// 第二个参数负数表示从右边开始删除几个等于100的元素
	client.LRem("key", -2, 100)

	// 如果存在多个100，第二个参数为0，表示删除所有元素等于100的数据
	client.LRem("key", 0, 100)

	//LIndex  根据索引坐标，查询列表中的数据
	// 列表索引从0开始计算，这里返回第6个元素
	val, err = client.LIndex("name", 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	//LInsert  在指定位置插入数据
	// 在列表中5的前面插入4
	// before是之前的意思
	err = client.LInsert("name", "before", 3, 88).Err()
	if err != nil {
		panic(err)
	}

	// 在列表中 luo 元素的前面插入 欢迎你
	client.LInsert("name", "before", "luo", "欢迎你")

	// 在列表中 luo 元素的后面插入 2019
	client.LInsert("name", "after", "luo", "2019")

}
