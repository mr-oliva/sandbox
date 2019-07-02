package cache

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/bookun/sandbox/go/judge-clientip/entity"
	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	client *memcache.Client
}

func NewMemcache(server ...string) *Memcache {
	mc := memcache.New(server...)
	return &Memcache{mc}
}

func (m *Memcache) Add(ctx context.Context, ip string, result entity.Result) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(result); err != nil {
		return err
	}
	return m.client.Set(&memcache.Item{Key: ip, Value: buf.Bytes()})
}

func (m *Memcache) Get(ctx context.Context, ip string) (entity.Result, error) {
	result := entity.Result{}
	item, err := m.client.Get(ip)
	if err != nil {
		return result, err
	}
	value := item.Value
	buf := bytes.NewBuffer(value)
	if err := json.NewDecoder(buf).Decode(&result); err != nil {
		return entity.Result{}, err
	}
	result.From = "cache"
	return result, nil
}
