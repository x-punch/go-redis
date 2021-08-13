/*
Package redis implements a simple cache management library base on redis
*/
package redis

import (
	"context"
	"strconv"
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

func (s *redisService) Subscribe(ctx context.Context, channels ...string) (*redis.PubSub, error) {
	pubsub := s.client.Subscribe(ctx, channels...)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

func (s *redisService) PSubscribe(ctx context.Context, channels ...string) (*redis.PubSub, error) {
	pubsub := s.client.PSubscribe(ctx, channels...)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

func (s *redisService) SubscribeExpired(ctx context.Context) (*redis.PubSub, error) {
	return s.Subscribe(ctx, "__keyevent@"+strconv.Itoa(s.config.DB)+"__:expired")
}
