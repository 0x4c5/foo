package example_service

import (
	"context"
	"foo/facades"
	"time"
)

type ExampleService struct{}

func New() *ExampleService {
	return &ExampleService{}
}

func (es *ExampleService) Put(key, val string) (err error) {
	// example service
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// set redis
	err = facades.Redis.Set(ctx, key, val, 0).Err()
	return
}

func (es *ExampleService) Get(key string) (val string, err error) {
	// do stuff here
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	val, err = facades.Redis.Get(ctx, key).Result()
	return
}
