package main

import (
	"context"
	"fmt"
	cache "github.com/go-pkgz/expirable-cache"
	"github.com/sfshf/sprout/app/govern/config"
	"time"
)

func NewCache(ctx context.Context) (cache.Cache, func(), error) {
	c := config.C.Cache
	var opts []cache.Option
	if c.IsLRU {
		opts = append(opts, cache.LRU())
	}
	if c.MaxKeys > 0 {
		opts = append(opts, cache.MaxKeys(c.MaxKeys))
	}
	if c.TTL > 0 {
		opts = append(opts, cache.TTL(time.Duration(c.TTL)*time.Minute))
	}
	opts = append(opts, cache.OnEvicted(func(key string, value interface{}) {
		fmt.Printf("Cache: %s was evicted.\n", key)
	}))
	cache, err := cache.NewCache(opts...)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	tick := time.Tick(time.Duration(c.TTL / 2))
	go func() {
		for {
			select {
			case <-tick:
				cache.DeleteExpired()
			case <-ctx.Done():
				return
			}
		}
	}()
	return cache, cancel, nil
}
