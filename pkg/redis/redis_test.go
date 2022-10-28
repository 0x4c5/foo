package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	rdb, err := Init("localhost:6379", "", 0, 100)
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err = rdb.Set(ctx, "key", "value", 0).Err()
	assert.NoError(t, err)
	val, err := rdb.Get(ctx, "key").Result()
	assert.NoError(t, err)
	t.Log(val)

}
