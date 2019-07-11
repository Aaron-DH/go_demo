package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"time"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:        viper.GetString("redis.addr"),
		Password:    viper.GetString("redis.password"),
		DB:          viper.GetInt("redis.db"),
		PoolSize:    viper.GetInt("redis.poolsize"), // Redis连接池大小
		MaxRetries:  3,
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Errorf(err, "Connect to redis error!")
	}
}

func Set(key string, value interface{}, expired time.Duration) error {
	switch value.(type) {
	case string:
		return client.Set(key, value, expired).Err()
	case int:
		return client.Set(key, value, expired).Err()
	default:
		jsonvalue, err := json.Marshal(&value)
		if err != nil {
			return err
		}
		return client.Set(key, jsonvalue, expired).Err()
	}
}

func Get(key string) (interface{}, error) {
	return client.Get(key).Result()
}
