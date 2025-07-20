package resources

import (
	"AuthTemplate/src"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

type Cache struct {
	Ctx    context.Context
	Client *redis.Client
}

func SetupRedis() {

	addr, err := redis.ParseURL(src.Config.RedisUrl)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: addr.Addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("failed to connect to Redis: %s", err))
		return
	}
	RedisClient = rdb
}

func (s *Cache) Set(key string, value interface{}, expiration time.Duration) {

	err := s.Client.Set(s.Ctx, key, value, expiration).Err()
	if err != nil {
		panic(fmt.Sprintf("failed to set key: %s", err))
	}
}
func (s *Cache) Get(key string) interface{} {
	val, err := s.Client.Get(s.Ctx, key).Result()
	if err != nil || val == "" {
		return nil
	}
	return val
}

func (s *Cache) Delete(key string) {
	s.Client.Del(s.Ctx, key)
	return
}

func (s *Cache) Exists(key string) bool {
	count, _ := s.Client.Exists(s.Ctx, key).Result()
	return count > 0
}

func (s *Cache) SetExpiration(key string, expiration time.Duration) {
	s.Client.Expire(s.Ctx, key, expiration)
}

func (s *Cache) GetTTL(key string) time.Duration {
	r, err := s.Client.TTL(s.Ctx, key).Result()
	if err != nil {
		panic(err)
		return 0
	}
	return r
}
