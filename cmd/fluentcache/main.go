package main

import (
	"time"

	"github.com/jkaveri/fluent-option/fluentcache"
)

func main() {
	connection := "localhost"
	password := "JA$2mAe%1@s5"
	db := 0
	readTimeout := 2 * time.Second
	writeTimeout := 2 * time.Second
	dialTimeout := 5 * time.Second
	masterName := "master"
	sentinelAddrs := []string{"localhost:6380", "localhost:6381", "localhost:6382"}
	sentinelPassword := ""
	maxRetries := 5
	minRetryBackoff := 10 * time.Second
	maxRetryBackoff := 5 * time.Minute

	redisCache := fluentcache.NewRedisCache(
		fluentcache.UseSentinelRedis(
			masterName,
			sentinelAddrs,
			sentinelPassword,
		),
		fluentcache.WithDB(db),
		fluentcache.WithPassword(password),
		fluentcache.WithTimeoutPolicy(
			dialTimeout,
			readTimeout,
			writeTimeout,
		),
		fluentcache.WithRetryPolicy(
			maxRetries,
			minRetryBackoff,
			maxRetryBackoff,
		),
	)

	// set key
	_ = redisCache.Set("my_key", "my_value")
}
