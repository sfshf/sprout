package cache

import cache "github.com/go-pkgz/expirable-cache"

type MemoryCache struct {
	cache cache.Cache
}

func NewMemoryCache(cache cache.Cache) *MemoryCache {
	return &MemoryCache{
		cache: cache,
	}
}

func (a *MemoryCache) Engine() cache.Cache {
	return a.cache
}
