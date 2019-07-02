package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/bookun/sandbox/go/judge-clientip/entity"
	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	client redis.Conn
}

func NewRedis(addr, pass string) (*Redis, error) {
	//redisPool := &redis.Pool{
	//	Dial: func() (redis.Conn, error) {
	//		conn, err := redis.Dial("tcp", addr)
	//		if pass == "" {
	//			return conn, err
	//		}
	//		if err != nil {
	//			return nil, err
	//		}
	//		if _, err := conn.Do("AUTH", pass); err != nil {
	//			conn.Close()
	//			return nil, err
	//		}
	//		return conn, nil
	//	},
	//}
	//return &Redis{redisPool}
	opt := redis.DialPassword(pass)
	conn, err := redis.Dial("tcp", addr, opt)
	if err != nil {
		return nil, err
	}
	return &Redis{conn}, nil
}

func (r *Redis) Add(ctx context.Context, ip string, result entity.Result) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(result); err != nil {
		return err
	}
	fmt.Println(buf.String())
	_, err := r.client.Do("SET", ip, buf.String())
	return err
}

func (r *Redis) Get(ctx context.Context, ip string) (entity.Result, error) {
	result := entity.Result{}

	exists, err := redis.Bool(r.client.Do("EXISTS", ip))
	if err != nil {
		return result, err
	}
	if !exists {
		return result, nil
	}

	cache, err := redis.String(r.client.Do("GET", ip))
	if err != nil {
		return result, err
	}
	buf := bytes.NewBufferString(cache)
	if err := json.NewDecoder(buf).Decode(&result); err != nil {
		return result, err
	}
	result.From = "cache"
	return result, nil
}
