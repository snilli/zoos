package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis(addr string, password string, db int, minidleconns int, poolsize int, connMaxIdleTime time.Duration, poolTimeout time.Duration) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:            addr,            // redis-host
		Password:        password,        // redis-password
		DB:              db,              // redis-DB
		MinIdleConns:    minidleconns,    // use default
		PoolSize:        poolsize,        //use default
		ConnMaxIdleTime: connMaxIdleTime, //use default
		PoolTimeout:     poolTimeout,     //use default
	})

}
