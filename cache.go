package internal

import (
	"context"
	"google.golang.org/appengine/v2/memcache"
)

func AddToCache(ctx context.Context, key string, value []byte) {
	item1 := &memcache.Item{
		Key:   key,
		Value: []byte(value),
	}
	if err := memcache.Set(ctx, item1); err != nil {
		panic(err)
	}
}

func GetFromCache(ctx context.Context, key string) ([]byte, bool) {
	item0, err := memcache.Get(ctx, key)
	if err != nil && err != memcache.ErrCacheMiss {
		return []byte(""), false
	} else if err == memcache.ErrCacheMiss {
		return []byte(""), false
	} else {
		return item0.Value, true
	}
}
