package cachev1

type RedisCache struct {
	connection string
	db         int
	password   string
}

// RedisCache create new redis cache
func NewRedisCache(connection, password string, db int) *RedisCache {
	return &RedisCache{
		connection: connection,
		password:   password,
		db:         db,
	}
}

func (*RedisCache) Set(key, value string) error {
	return nil
}

func (*RedisCache) Get(key string) (string, error) {
	return "", nil
}
