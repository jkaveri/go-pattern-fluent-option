package cachev3

import "time"

type RedisCache struct {
	connection string
	db         int
	password   string

	maxRetries      int
	minRetryBackoff time.Duration
	maxRetryBackoff time.Duration

	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// RedisCache create new redis cache
func NewRedisCache(connection, password string, db int) *RedisCache {
	return &RedisCache{
		connection: connection,
		password:   password,
		db:         db,
	}
}

func NewRedisCacheWithRetryPolicy(
	connection, password string,
	db, maxRetries int,
	minRetryBackoff, maxRetryBackoff time.Duration,
) *RedisCache {
	rc := NewRedisCache(connection, password, db)

	rc.maxRetries = maxRetries
	rc.minRetryBackoff = minRetryBackoff
	rc.maxRetryBackoff = maxRetryBackoff
	return rc
}

func NewRedisCacheWithTimeoutPolicy(
	connection, password string,
	db, maxRetries int,
	dialTimeout, readTimeout, writeTimeout time.Duration,
) *RedisCache {
	rc := NewRedisCache(connection, password, db)
	rc.dialTimeout = dialTimeout
	rc.readTimeout = readTimeout
	rc.writeTimeout = writeTimeout
	return rc
}

func (*RedisCache) Set(key, value string) error {
	return nil
}

func (*RedisCache) Get(key string) (string, error) {
	return "", nil
}
