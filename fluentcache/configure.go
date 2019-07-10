package fluentcache

import (
	"crypto/tls"
	"time"
)

func UseStandalone(connection string) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.sentinel = false
		rc.connection = connection
	}
}

func UseSentinelRedis(
	masterName string,
	sentinelAddrs []string,
	sentinelPassword string,
) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.sentinel = true
		rc.masterName = masterName
		rc.sentinelAddrs = sentinelAddrs
		rc.sentinelPassword = sentinelPassword
	}
}

func WithPassword(password string) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.password = password
	}
}

func WithDB(db int) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.db = db
	}
}

func WithDialer(dialer DialerFunc) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.dialer = dialer
	}
}

func OnConnect(onConnectFunc OnConnectFunc) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.onConnect = onConnectFunc
	}
}

func WithRetryPolicy(
	maxRetries int,
	minRetryBackoff, maxRetryBackoff time.Duration,
) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.maxRetries = maxRetries
		rc.minRetryBackoff = minRetryBackoff
		rc.maxRetryBackoff = maxRetryBackoff
	}
}

func WithTimeoutPolicy(
	dialTimeout, readTimeout, writeTimeout time.Duration,
) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.dialTimeout = dialTimeout
		rc.readTimeout = readTimeout
		rc.writeTimeout = writeTimeout
	}
}

func WithConnectionPoolPolicy(
	poolSize, minIdleConns int,
	maxConnAge, poolTimeout, idleTimeout, idleCheckFrequency time.Duration,
) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.poolSize = poolSize
		rc.minIdleConns = minIdleConns
		rc.maxConnAge = maxConnAge
		rc.poolTimeout = poolTimeout
		rc.idleTimeout = idleTimeout
		rc.idleCheckFrequency = idleCheckFrequency
	}
}

func WithTLSOptions(tlsConfig *tls.Config) ConfigureFunc {
	return func(rc *RedisCache) {
		rc.tLsConfig = tlsConfig
	}
}
