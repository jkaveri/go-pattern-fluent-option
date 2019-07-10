## Problem
When design API in Golang I usually create some contractor function that helps to configure the service. Like below

```go
// cache/cache.go
package cache

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
```
*I skipped some implementation detail to keep the code simple, you only need to focus the constructor function*

Then I can init my `RedisCache` like this:
```go
// main.go
package main

import "github.com/jkaveri/fluent-option/cache"

func main() {
	connection := "localhost"
	password := "JA$2mAe%1@s5"
	db := 0

	redisCache := cache.NewRedisCache(connection, password, 0)

	// set key
	_ = redisCache.Set("my_key", "my_value")
}
```

The constructor looks good enough and I started to use the constructor function in many places. However, I got a problem when I want to add more argument into the constructor function, the problems are:

* In Go, we don't have the **optional argument**
* In Go, we don't have **Overloading**  

So we only have these options:

* Add new arguments into existing function and update existing code. This option is worse because of that impact on the existing code. Furthermore, in case your API is the dependency of other packages and this option can is a breaking change

    ```go
    // cachev2.go
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
    ```

* Create new function like this `NewFailoverRedis(masterName string, sentinelAddrs []string, sentinelPassword string, maxRetries, maxBackoff string)`, creating a new function makes sense but what if we will need more arguments in future? Of course, you can say we need compliance with YAGNI, but sometimes if we have an option that more flexible so that the API more stable. In other words, in a future version of API, the API's clients don't struggle about the breaking changes

    ```go
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
    ```