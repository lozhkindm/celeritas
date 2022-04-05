package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int) error
	Forget(string) error
	Empty() error
	EmptyByMatch(string) error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]interface{}

func (rc *RedisCache) Has(key string) (bool, error) {
	key = fmt.Sprintf("%s:%s", rc.Prefix, key)
	conn := rc.Conn.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)

	found, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return found, err
}

func (rc *RedisCache) Get(key string) (interface{}, error) {
	return nil, nil
}

func (rc *RedisCache) Set(key string, val interface{}, ttl ...int) error {
	return nil
}

func (rc *RedisCache) Forget(key string) error {
	return nil
}

func (rc *RedisCache) Empty() error {
	return nil
}

func (rc *RedisCache) EmptyByMatch(pattern string) error {
	return nil
}
