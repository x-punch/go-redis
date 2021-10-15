package redis

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func (s *redisService) Subscribe(ctx context.Context, channels ...string) (*redis.PubSub, error) {
	pubsub := s.client.Subscribe(ctx, channels...)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, err
	}
	return pubsub, nil
}

func (s *redisService) PubsubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return s.client.PubSubChannels(ctx, pattern)
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
