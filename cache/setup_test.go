package cache

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/dgraph-io/badger/v3"
	"github.com/gomodule/redigo/redis"
)

var (
	testRedisCache  RedisCache
	testBadgerCache BadgerCache
)

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

	_ = os.RemoveAll("./testdata/tmp/badger")

	if _, err := os.Stat("./testdata/tmp"); os.IsNotExist(err) {
		if err := os.Mkdir("./testdata/tmp", 0755); err != nil {
			log.Fatalf("could not create a tmp folder: %s", err)
		}
	}
	if err := os.Mkdir("./testdata/tmp/badger", 0755); err != nil {
		log.Fatalf("could not create a badger folder: %s", err)
	}

	testBadgerCache.Conn, err = badger.Open(badger.DefaultOptions("./testdata/tmp/badger"))
	if err != nil {
		log.Fatalf("could not create a badger db: %s", err)
	}

	os.Exit(m.Run())
}
