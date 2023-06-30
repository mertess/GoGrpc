package redisClient

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisAddr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})

	return &RedisClient{client: client}
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.client.Get(key).Result()
}

func (rc *RedisClient) Set(key string, value any) error {
	return rc.client.Set(key, value, 5*time.Second).Err()
}
