package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func (s *redisService) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return s.client.Eval(ctx, script, keys, args...)
}

func (s *redisService) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return s.client.EvalSha(ctx, sha1, keys, args...)
}

func (s *redisService) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return s.client.ScriptExists(ctx, hashes...)
}

func (s *redisService) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	return s.client.ScriptFlush(ctx)
}

func (s *redisService) ScriptKill(ctx context.Context) *redis.StatusCmd {
	return s.client.ScriptKill(ctx)
}

func (s *redisService) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return s.client.ScriptLoad(ctx, script)
}
