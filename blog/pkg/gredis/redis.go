package gredis

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"

	"lianxi/blog/pkg/setting"
)

var RedisConn *redis.Pool

// 初始化Redis实例
func Setup() error {
	/*
		Pool维护一个连接池。应用程序调用Get方法从池中获取连接，
		并调用连接的Close方法将连接的资源返回到池中。
	*/
	RedisConn = &redis.Pool{
		//最大空闲连接数。
		MaxIdle: setting.RedisSetting.MaxIdle,
		//在给定时间内分配的最大连接数。当为0时，池中的连接数没有限制。
		MaxActive: setting.RedisSetting.MaxActive,
		//在此期间保持空闲状态后关闭连接。如果该值为零，则不关闭空闲连接。
		//应用程序应该将超时设置为一个小于服务器超时的值。
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		/*
			Dial是应用程序提供的用于创建和配置连接的功能。
			从Dial返回的连接不能处于特殊状态(订阅pubsub通道、事务启动……)。
		*/
		Dial: func() (redis.Conn, error) {
			//拨号连接到Redis服务器在给定的网络和地址使用指定的选项。
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				//Do向服务器发送命令并返回接收到的应答。
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					//Close关闭连接。
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		/*
			TestOnBorrow是一个可选的应用程序提供的函数，
			用于在应用程序再次使用空闲连接之前检查该连接的健康状况。
			参数t是连接被返回到池的时间。如果函数返回错误，则连接关闭。
		*/
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// Set设置一个键/值
func Set(key string, data interface{}, time int) error {
	/*
		Get得到连接。应用程序必须关闭返回的连接。此方法总是返回一个
		有效的连接，以便应用程序可以将错误处理推迟到第一次使用该连接时。
		如果在获取底层连接时出现错误，
		则连接Err、Do、Send、Flush和Receive方法将返回该错误。
	*/
	conn := RedisConn.Get()
	defer conn.Close()
	//Marshal返回值的JSON编码。
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//设置key和value
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	//设置到期时间
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists 检查键是否存在
func Exists(key string) bool {
	//Get得到连接。应用程序必须关闭返回的连接
	conn := RedisConn.Get()
	defer conn.Close()
	//Bool是一个助手，它将命令回复转换为布尔值。如果err不等于nil,
	//Bool返回false, err。否则Bool将回复转换为布尔值
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get 获取一个key
func Get(key string) ([]byte, error) {
	//获取一个连接，应用程序必须关闭返回的连接
	conn := RedisConn.Get()
	defer conn.Close()

	/*
		Bytes是一个助手，它将命令回复转换为字节片。如果err不等于nil，
		则Bytes返回nil, err。否则Bytes将回复转换为一个字节片
	*/
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete 删除一个key
func Delete(key string) (bool, error) {
	//获取一个连接，应用程序必须关闭返回的连接
	conn := RedisConn.Get()
	defer conn.Close()
	//Bool是一个助手，它将命令回复转换为布尔值。如果err不等于nil,
	//Bool返回false, err。否则Bool将回复转换为布尔值
	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes批量删除
func LikeDeletes(key string) error {
	//获取一个连接，应用程序必须关闭返回的连接
	conn := RedisConn.Get()
	defer conn.Close()
	/*
		string是一个助手，它将数组命令回复转换为[]字符串。
		如果err不等于nil，那么string返回nil, err。
		在输出片中Nil数组项被转换为""。如果数组项不是批量字符串或nil，
		则string返回错误。
	*/
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	//遍历[]切片，删除key
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
