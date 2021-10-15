/*
Package redis implements a simple cache management library base on redis
*/
package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// ErrKeyNotExist reply Redis returns when key does not exist.
const ErrKeyNotExist = redis.Nil

type redisService struct {
	config Config
	client *redis.Client
}

// NewService create a new cache service to access redis
func NewService(config Config) Service {
	client := redis.NewClient(&redis.Options{
		Network:  config.Network,
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	return &redisService{config, client}
}

func (s *redisService) Client() *redis.Client {
	return s.client
}

func (s *redisService) Del(ctx context.Context, key string) error {
	return s.client.Del(ctx, key).Err()
}

func (s *redisService) Exist(ctx context.Context, key string) (bool, error) {
	count, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *redisService) Exists(ctx context.Context, keys ...string) (int64, error) {
	return s.client.Exists(ctx, keys...).Result()
}

func (s *redisService) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return s.client.Expire(ctx, key, expiration).Result()
}

func (s *redisService) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return s.client.ExpireAt(ctx, key, tm).Result()
}

func (s *redisService) TTL(ctx context.Context, key string) (time.Duration, error) {
	return s.client.TTL(ctx, key).Result()
}

func (s *redisService) Keys(ctx context.Context, pattern string) ([]string, error) {
	return s.client.Keys(ctx, pattern).Result()
}

func (s *redisService) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return s.client.Scan(ctx, cursor, match, count).Result()
}

func (s *redisService) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	return s.client.Do(ctx, args...)
}
