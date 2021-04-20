package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	cache "github.com/go-pkgz/expirable-cache"
	"github.com/go-redis/redis/v8"
	"github.com/sfshf/sprout/app/govern/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context) (*mongo.Database, func(), error) {
	c := config.C.MongoDB
	srvUri, err := url.Parse(c.ServerUri)
	if err != nil {
		return nil, nil, err
	}
	cliOpt := options.Client().SetHosts([]string{srvUri.Host})
	if direct := srvUri.Query().Get("directConnection"); direct != "" && strings.ToUpper(direct) == "TRUE" {
		cliOpt.SetDirect(true)
	}
	if dbName := srvUri.Path[1:]; dbName != "" {
		c.Database = dbName
	}
	client, err := mongo.NewClient(cliOpt)
	if err != nil {
		return nil, nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, nil, err
	}
	log.Println("Mongo DB is on!!!")
	return client.Database(c.Database), func() {
		client.Disconnect(ctx)
	}, nil
}

func NewRedisDB(ctx context.Context) (*redis.Client, func(), error) {
	c := config.C.Redis
	if !c.Enable {
		return nil, nil, nil
	}
	cli := redis.NewClient(&redis.Options{
		Network:  c.Network,
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB,
	})
	if _, err := cli.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	log.Println("Redis DB is on!!!")
	return cli, func() {
		cli.Close()
	}, nil
}

func NewCache(ctx context.Context) (cache.Cache, func(), error) {
	c := config.C.Cache
	if !c.Enable {
		return nil, nil, nil
	}
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
	log.Println("Cache DB is on!!!")
	return cache, cancel, nil
}
