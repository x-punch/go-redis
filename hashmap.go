package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func (s *redisService) HGet(ctx context.Context, key string, field string) (string, error) {
	str, err := s.client.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", ErrKeyNotExist
	}
	return str, err
}

func (s *redisService) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := s.client.HGetAll(ctx, key).Result()
	return result, err
}

func (s *redisService) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return s.client.HMGet(ctx, key, fields...).Result()
}

func (s *redisService) HExists(ctx context.Context, key, field string) (bool, error) {
	return s.client.HExists(ctx, key, field).Result()
}

func (s *redisService) HSet(ctx context.Context, key string, values ...interface{}) error {
	return s.client.HSet(ctx, key, values...).Err()
}

func (s *redisService) HDel(ctx context.Context, key string, fields ...string) error {
	return s.client.HDel(ctx, key, fields...).Err()
}

func (s *redisService) HIncrBy(ctx context.Context, key, field string, incr int64) error {
	return s.client.HIncrBy(ctx, key, field, incr).Err()
}

func (s *redisService) HIncrByFloat(ctx context.Context, key, field string, incr float64) error {
	return s.client.HIncrByFloat(ctx, key, field, incr).Err()
}

func (s *redisService) HKeys(ctx context.Context, key string) ([]string, error) {
	return s.client.HKeys(ctx, key).Result()
}

func (s *redisService) HLen(ctx context.Context, key string) (int64, error) {
	return s.client.HLen(ctx, key).Result()
}
