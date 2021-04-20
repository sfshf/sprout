package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

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

func (a *RedisCache) TokenExists(ctx context.Context, key string, value string) bool {
	res := a.cli.Get(ctx, key).Val()
	fmt.Println(res)
	return res == value
}

func (a *RedisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) bool {
	_, err := a.cli.Set(ctx, key, value, ttl).Result()
	return err == nil
}

func (a *RedisCache) Del(ctx context.Context, key string) bool {
	del := a.cli.Del(ctx, key).Val()
	return del == 1
}
