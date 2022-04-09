package cache

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
)

var testRedisCache RedisCache

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		log.Fatalf("cound not run a miniredis: %s", err)
	}
	defer s.Close()

	pool := redis.Pool{
		MaxIdle:     50,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", s.Addr())
		},
	}

	testRedisCache.Conn = &pool
	testRedisCache.Prefix = "test-celeritas"

	defer func(Conn *redis.Pool) {
		_ = Conn.Close()
	}(testRedisCache.Conn)

	os.Exit(m.Run())
}
