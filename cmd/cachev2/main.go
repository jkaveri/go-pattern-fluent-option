package main

import (
	"time"

	"github.com/jkaveri/fluent-option/cachev2"
)

func main() {
	connection := "localhost"
	password := "JA$2mAe%1@s5"
	db := 0
	maxRetries := 2
	minRetryWithBackoff := 2 * time.Second
	maxRetryWithBackoff := 2 * time.Second

	redisCache := cachev2.NewRedisCache(
		connection,
		password,
		db,
		maxRetries,
		minRetryWithBackoff,
		maxRetryWithBackoff,
	)

	// set key
	_ = redisCache.Set("my_key", "my_value")
}
