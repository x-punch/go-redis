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

// redisService represents cache service build on redis
type redisService struct {
	config *Config
	client *redis.Client
}

// NewService create a new cache service to access redis
func NewService(config *Config) Service {
	client := redis.NewClient(&redis.Options{
		Network:  config.Network,
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	return &redisService{config, client}
}

// GetBytes search given key in cache then return received reply
func (s *redisService) GetBytes(key string) ([]byte, error) {
	bytes, err := s.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, ErrKeyNotExist
	}
	return bytes, err
}

// FindBytes search given key in cache then return received reply
func (s *redisService) FindBytes(key string) ([]byte, error) {
	bytes, err := s.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return bytes, err
}

// GetString search given key in cache then return reply in string
func (s *redisService) GetString(key string) (string, error) {
	str, err := s.client.Get(key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotExist
	}
	return str, err
}

// FindString search given key in cache then return reply in string
func (s *redisService) FindString(key string) (string, error) {
	str, err := s.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return str, err
}

// Set given key with value in cache within given expired time
func (s *redisService) Set(key string, value []byte, expiration time.Duration) error {
	return s.client.Set(key, value, expiration).Err()
}

// Delete given key in cache
func (s *redisService) Delete(key string) error {
	return s.client.Del(key).Err()
}

// Exists check given key whether exist in cache
func (s *redisService) Exists(key string) (bool, error) {
	_, err := s.client.Get(key).Result()
	if err == redis.Nil {
		return true, nil
	}
	return false, err
}

// Subscribe the client to the specified channels. It returns
// empty subscription if there are no channels.
func (s *redisService) Subscribe(channels ...string) (*redis.PubSub, error) {
	pubsub := s.client.Subscribe(channels...)
	_, err := pubsub.Receive()
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

// PSubscribe the client to the given patterns. It returns
// empty subscription if there are no patterns.
func (s *redisService) PSubscribe(patterns ...string) (*redis.PubSub, error) {
	pubsub := s.client.PSubscribe(patterns...)
	_, err := pubsub.Receive()
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

func (s *redisService) SubscribeExpired() (*redis.PubSub, error) {
	return s.Subscribe("__keyevent@" + strconv.Itoa(s.config.DB) + "__:expired")
}
