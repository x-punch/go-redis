package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (s *redisService) GetBytes(ctx context.Context, key string) ([]byte, error) {
	bytes, err := s.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, ErrKeyNotExist
	}
	return bytes, err
}

func (s *redisService) GetString(ctx context.Context, key string) (string, error) {
	str, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotExist
	}
	return str, err
}

func (s *redisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.client.Set(ctx, key, value, expiration).Err()
}

func (s *redisService) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return s.client.SetNX(ctx, key, value, expiration).Result()
}

func (s *redisService) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return s.client.SetXX(ctx, key, value, expiration).Result()
}

func (s *redisService) SetRange(ctx context.Context, key string, offset int64, value string) error {
	return s.client.SetRange(ctx, key, offset, value).Err()
}

func (s *redisService) StrLen(ctx context.Context, key string) (int64, error) {
	return s.client.StrLen(ctx, key).Result()
}

func (s *redisService) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	return s.client.GetBit(ctx, key, offset).Result()
}

func (s *redisService) SetBit(ctx context.Context, key string, offset int64, value int) error {
	return s.client.SetBit(ctx, key, offset, value).Err()
}
