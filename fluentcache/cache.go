package fluentcache

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/go-redis/redis"
)

type ConfigureFunc = func(redisCache *RedisCache)

type DialerFunc = func(
	ctx context.Context,
	network,
	addr string,
) (net.Conn, error)

type OnConnectFunc = func(conn *redis.Conn) error

type RedisCache struct {
	sentinel   bool
	connection string
	db         int
	password   string

	sentinelAddrs    []string
	masterName       string
	sentinelPassword string

	dialer          DialerFunc
	onConnect       OnConnectFunc
	maxRetries      int
	minRetryBackoff time.Duration
	maxRetryBackoff time.Duration

	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration

	poolSize           int
	minIdleConns       int
	maxConnAge         time.Duration
	poolTimeout        time.Duration
	idleTimeout        time.Duration
	idleCheckFrequency time.Duration

	tLsConfig *tls.Config
}

// RedisCache create new redis cache
func NewRedisCache(configures ...ConfigureFunc) *RedisCache {
	var rc RedisCache
	for _, configure := range configures {
		configure(&rc)
	}
	return &rc
}

func (*RedisCache) Set(key, value string) error {
	return nil
}

func (*RedisCache) Get(key string) (string, error) {
	return "", nil
}
