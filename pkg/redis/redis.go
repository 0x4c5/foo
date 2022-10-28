package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func Init(addr, password string, db, poolSize int) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return
}
