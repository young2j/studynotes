package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func main() {
	options, err := redis.ParseURL("redis://username:password@localhost:6379/1")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(options)
	defer rdb.Close()

	rdb.Ping(context.Background())
}
