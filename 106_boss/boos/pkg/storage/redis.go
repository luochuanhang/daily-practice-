package storage

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

// RedisConfig stands for database connection configuration.
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func (p *RedisConfig) String() string {
	return fmt.Sprintf("redis://:%s@%s/%d",
		p.Password, p.Addr, p.DB)
}

type redisInstance struct {
	client *redis.Client
}

var (
	defaultRedis = &redisInstance{}
	onceRedis    sync.Once
)

// var _ Storage = &redisInstance{} // TODO uncomment this when using gorm@v2

func (p *redisInstance) Name() string {
	return "redis"
}

// Init init redis DB connection.
func Init(addr, password string, db int) error {
	var err error
	onceRedis.Do(func() {
		if defaultRedis.client == nil {
			defaultRedis.client, err = New(addr, password, db)
		}
	})
	if err != nil {
		return fmt.Errorf("redisInstance.New %w", err)
	}
	return nil
}

// Get get default redis connection.
func Get() *redis.Client { // TODO make this as a function of redisInstance when using gorm@v2.
	return defaultRedis.client
}

// New create a new redis connection.
func New(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
