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

	// Expire Set a key's time to live in seconds
	Expire(key string, expiration time.Duration) (bool, error)

	// ExpireAt Set the expiration for a key as a UNIX timestamp
	ExpireAt(key string, expiration time.Time) (bool, error)

	// Delete Delete a key
	Delete(key string) error

	// Exists Determine if a key exists
	Exists(key string) (bool, error)

	//Keys Find all keys matching the given pattern
	Keys(pattern string) ([]string, error)

	// Scan Incrementally iterate the keys space
	Scan(cursor uint64, match string, count int64) ([]string, uint64, error)

	// Subscribe Listen for messages published to the given channels
	Subscribe(channels ...string) (*redis.PubSub, error)

	// PSubscribe Listen for messages published to channels matching the given patterns
	PSubscribe(patterns ...string) (*redis.PubSub, error)

	// SubscribeExpired the client to the specified channels. It returns
	// empty subscription if there are no channels.
	SubscribeExpired() (*redis.PubSub, error)
}
