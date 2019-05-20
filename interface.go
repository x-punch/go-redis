package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// Service reprsents redis cache service
type Service interface {
	// GetBytes search given key in cache then return received reply
	GetBytes(key string) (reply []byte, err error)

	// FindBytes search given key in cache then return received reply
	FindBytes(key string) (reply []byte, err error)

	// GetString search given key in cache then return reply in string
	GetString(key string) (reply string, err error)

	// FindString search given key in cache then return reply in string
	FindString(key string) (reply string, err error)

	// Set given key with value in cache within given expired time
	Set(key string, value []byte, expiration time.Duration) error

	// Delete given key in cache
	Delete(key string) error

	// Exists check given key whether exist in cache
	Exists(key string) (bool, error)

	// Subscribe the client to the specified channels. It returns
	// empty subscription if there are no channels.
	Subscribe(channels ...string) (*redis.PubSub, error)

	// PSubscribe the client to the given patterns. It returns
	// empty subscription if there are no patterns.
	PSubscribe(patterns ...string) (*redis.PubSub, error)

	// SubscribeExpired the client to the specified channels. It returns
	// empty subscription if there are no channels.
	SubscribeExpired() (*redis.PubSub, error)
}
