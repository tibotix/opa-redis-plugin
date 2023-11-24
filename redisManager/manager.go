package redisManager

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/tibotix/opa-redis-plugin/proxy"
)

type RedisManager struct {
	RedisProxy   *proxy.Proxy[*redis.Client]
	RedisContext context.Context
}

func New() *RedisManager {
	redisContext := context.Background()
	redisProxy := &proxy.Proxy[*redis.Client]{}

	return &RedisManager{
		RedisContext: redisContext,
		RedisProxy:   redisProxy,
	}
}

func (m *RedisManager) RegisterCommands() {
	m.registerAutogenCommands()
	m.registerManualCommands()
}
