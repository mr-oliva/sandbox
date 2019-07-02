package cache

import (
	"context"
	"time"

	"google.golang.org/appengine/memcache"
)

type GoogleMemcache struct {
}

func (m *GoogleMemcache) Add(ctx context.Context, ip, result string) error {
	item := &memcache.Item{
		Key:        ip,
		Value:      []byte(result),
		Expiration: time.Duration(7*24) * time.Hour,
	}
	err := memcache.Add(ctx, item)
	if err != nil && err != memcache.ErrNotStored {
		return err
	}
	return nil
}

func (m *GoogleMemcache) Get(ctx context.Context, ip string) (string, error) {
	item, err := memcache.Get(ctx, ip)
	if err != nil && err != memcache.ErrCacheMiss {
		return "", err
	}
	if err == memcache.ErrCacheMiss {
		return "", nil
	}
	return string(item.Value), nil
}
