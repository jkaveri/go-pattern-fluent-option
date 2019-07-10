package main

import "github.com/jkaveri/fluent-option/cache"

func main() {
	connection := "localhost"
	password := "JA$2mAe%1@s5"
	db := 0

	redisCache := cache.NewRedisCache(connection, password, db)

	// set key
	_ = redisCache.Set("my_key", "my_value")
}
