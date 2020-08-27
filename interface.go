package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Service reprsents redis cache service
type Service interface {
	// Client will return redis client
	Client() *redis.Client

	// GetBytes search given key in cache then return received reply
	GetBytes(ctx context.Context, key string) (reply []byte, err error)

	// FindBytes search given key in cache then return received reply
	FindBytes(ctx context.Context, key string) (reply []byte, err error)

	// GetString search given key in cache then return reply in string
	GetString(ctx context.Context, key string) (reply string, err error)

	// FindString search given key in cache then return reply in string
	FindString(ctx context.Context, key string) (reply string, err error)

	// Set given key with value in cache within given expired time
	Set(ctx context.Context, key string, value []byte, expiration time.Duration) error

	// Expire Set a key's time to live in seconds
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)

	// ExpireAt Set the expiration for a key as a UNIX timestamp
	ExpireAt(ctx context.Context, key string, expiration time.Time) (bool, error)

	// Delete Delete a key
	Delete(ctx context.Context, key string) error

	// Exists Determine if a key exists
	Exists(ctx context.Context, key string) (bool, error)

	//Keys Find all keys matching the given pattern
	Keys(ctx context.Context, pattern string) ([]string, error)

	// Scan Incrementally iterate the keys space
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)

	// Subscribe Listen for messages published to the given channels
	Subscribe(ctx context.Context, channels ...string) (*redis.PubSub, error)

	// PSubscribe Listen for messages published to channels matching the given patterns
	PSubscribe(ctx context.Context, patterns ...string) (*redis.PubSub, error)

	// SubscribeExpired the client to the specified channels. It returns
	// empty subscription if there are no channels.
	SubscribeExpired(ctx context.Context) (*redis.PubSub, error)
}
