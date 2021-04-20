package cache

import "github.com/go-redis/redis/v8"

type RedisCache struct {
	cli *redis.Client
}

func NewRedisCache(cli *redis.Client) *RedisCache {
	return &RedisCache{
		cli: cli,
	}
}

func (a *RedisCache) Engine() *redis.Client {
	return a.cli
}
