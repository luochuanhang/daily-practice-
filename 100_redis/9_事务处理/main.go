package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.18.107.135:6379", //redis地址
		Password: "",                    //redis密码，没有则为空
		DB:       0,                     //默认数据库，默认是0
	})

	//1.TxPipeline
	//以Pipeline的方式操作事务
	pipe := client.TxPipeline()
	// 执行事务操作，可以通过pipe读写redis
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	// 上面代码等同于执行下面redis命令
	//
	//     MULTI
	//     INCR pipeline_counter
	//     EXPIRE pipeline_counts 3600
	//     EXEC

	// 通过Exec函数提交redis事务
	_, err := pipe.Exec()

	// 提交事务后，我们可以查询事务操作的结果
	// 前面执行Incr函数，在没有执行exec函数之前，实际上还没开始运行。
	fmt.Println(incr.Val(), err)

	//2.watch
	//redis乐观锁支持，可以通过watch监听一些Key, 如果这些key的值没有被其他人改变的话，才可以提交事务。
	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get("key").Result()
		if err != nil && err != redis.Nil {
			return err
		}

		// 这里可以处理业务
		fmt.Println(v)

		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set("key", "new value", 0)
			return nil
		})
		return err
	}

	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
	client.Watch(fn, "key")

}
