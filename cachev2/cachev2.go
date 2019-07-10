package cachev2

import "time"

type RedisCache struct {
	connection string
	db         int
	password   string

	maxRetries      int
	minRetryBackoff time.Duration
	maxRetryBackoff time.Duration
}

// RedisCache create new redis cache
func NewRedisCache(
	connection, password string,
	db, maxRetries int,
	minRetryBackoff, maxRetryBackoff time.Duration,
) *RedisCache {

	return &RedisCache{
		connection:      connection,
		password:        password,
		db:              db,
		maxRetries:      maxRetries,
		minRetryBackoff: minRetryBackoff,
		maxRetryBackoff: maxRetryBackoff,
	}
}

func (*RedisCache) Set(key, value string) error {
	return nil
}

func (*RedisCache) Get(key string) (string, error) {
	return "", nil
}
