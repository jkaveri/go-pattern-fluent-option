package main

import "github.com/jkaveri/fluent-option/cachev2"

func main() {
	connection := "localhost"
	password := "JA$2mAe%1@s5"
	db := 0

	redisCache := cachev2.NewRedisCache(connection, password, db)

	// set key
	_ = redisCache.Set("my_key", "my_value")
}
