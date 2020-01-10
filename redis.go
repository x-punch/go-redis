/*
Package redis implements a simple cache management library base on redis
*/
package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
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

func (s *redisService) GetBytes(key string) ([]byte, error) {
	bytes, err := s.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, ErrKeyNotExist
	}
	return bytes, err
}

func (s *redisService) FindBytes(key string) ([]byte, error) {
	bytes, err := s.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return bytes, err
}

func (s *redisService) GetString(key string) (string, error) {
	str, err := s.client.Get(key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotExist
	}
	return str, err
}

func (s *redisService) FindString(key string) (string, error) {
	str, err := s.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return str, err
}

func (s *redisService) Set(key string, value []byte, expiration time.Duration) error {
	return s.client.Set(key, value, expiration).Err()
}

func (s *redisService) Delete(key string) error {
	return s.client.Del(key).Err()
}

func (s *redisService) Exists(key string) (bool, error) {
	_, err := s.client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *redisService) Expire(key string, expiration time.Duration) (bool, error) {
	return s.client.Expire(key, expiration).Result()
}

func (s *redisService) ExpireAt(key string, tm time.Time) (bool, error) {
	return s.client.ExpireAt(key, tm).Result()
}

func (s *redisService) Keys(pattern string) ([]string, error) {
	return s.client.Keys(pattern).Result()
}

func (s *redisService) Scan(cursor uint64, match string, count int64) ([]string, error) {
	iter := s.client.Scan(cursor, match, count).Iterator()
	result := []string{}
	for iter.Next() {
		result = append(result, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *redisService) Subscribe(channels ...string) (*redis.PubSub, error) {
	pubsub := s.client.Subscribe(channels...)
	_, err := pubsub.Receive()
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

func (s *redisService) PSubscribe(channels ...string) (*redis.PubSub, error) {
	pubsub := s.client.PSubscribe(channels...)
	_, err := pubsub.Receive()
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

func (s *redisService) SubscribeExpired() (*redis.PubSub, error) {
	return s.Subscribe("__keyevent@" + strconv.Itoa(s.config.DB) + "__:expired")
}
