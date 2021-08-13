package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Service reprsents redis cache service
type Service interface {
	// GetBytes search given key in cache then return received reply
	GetBytes(ctx context.Context, key string) (reply []byte, err error)

	// GetString search given key in cache then return reply in string
	GetString(ctx context.Context, key string) (reply string, err error)

	// Redis `SET key value [expiration]` command.
	//
	// Use expiration for `SETEX`-like behavior.
	// Zero expiration means the key has no expiration time.
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Redis `SET key value [expiration] NX` command.
	//
	// Zero expiration means the key has no expiration time.
	SetNX(context.Context, string, interface{}, time.Duration) (bool, error)

	// Redis `SET key value [expiration] XX` command.
	//
	// Zero expiration means the key has no expiration time.
	SetXX(context.Context, string, interface{}, time.Duration) (bool, error)

	SetRange(ctx context.Context, key string, offset int64, value string) error

	StrLen(ctx context.Context, key string) (int64, error)

	GetBit(ctx context.Context, key string, offset int64) (int64, error)

	SetBit(ctx context.Context, key string, offset int64, value int) error

	// Expire Set a key's time to live in seconds
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)

	// ExpireAt Set the expiration for a key as a UNIX timestamp
	ExpireAt(ctx context.Context, key string, expiration time.Time) (bool, error)

	// Del Delete a key
	Del(ctx context.Context, key string) error

	// Exist Determine if a key exists
	Exist(ctx context.Context, key string) (bool, error)

	// Exist Determine if some keys exists
	Exists(ctx context.Context, keys ...string) (int64, error)

	TTL(ctx context.Context, key string) (time.Duration, error)

	//Keys Find all keys matching the given pattern
	Keys(ctx context.Context, pattern string) ([]string, error)

	// Scan Incrementally iterate the keys space
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)

	// HMGet returns the values for the specified fields in the hash stored at key.
	HGet(ctx context.Context, key string, field string) (string, error)

	// HMGet returns the values for the specified fields in the hash stored at key.
	// It returns an interface{} to distinguish between empty string and nil value.
	HMGet(context.Context, string, ...string) ([]interface{}, error)

	// HSet accepts values in following formats:
	//   - HSet("myhash", "key1", "value1", "key2", "value2")
	//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
	//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
	//
	// Note that it requires Redis v4 for multiple field/value pairs support.
	HSet(context.Context, string, ...interface{}) error

	HGetAll(ctx context.Context, key string) (map[string]string, error)

	HExists(ctx context.Context, key, field string) (bool, error)

	HDel(ctx context.Context, key string, fields ...string) error

	HIncrBy(ctx context.Context, key, field string, incr int64) error

	HIncrByFloat(ctx context.Context, key, field string, incr float64) error

	HKeys(ctx context.Context, key string) ([]string, error)

	HLen(ctx context.Context, key string) (int64, error)

	// Subscribe Listen for messages published to the given channels
	Subscribe(ctx context.Context, channels ...string) (*redis.PubSub, error)

	// PSubscribe Listen for messages published to channels matching the given patterns
	PSubscribe(ctx context.Context, patterns ...string) (*redis.PubSub, error)

	// SubscribeExpired the client to the specified channels. It returns
	// empty subscription if there are no channels.
	SubscribeExpired(ctx context.Context) (*redis.PubSub, error)
}
