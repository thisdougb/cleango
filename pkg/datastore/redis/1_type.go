package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Datastore struct {
	host     string
	port     string
	database string
	ctx      context.Context
	client   *redis.Client
}

func NewRedisDatastore(host string, port string) *Datastore {

	ds := &Datastore{host, port, "0", context.Background(), nil}
	if ds.Connect() {
		return ds
	}
	return nil
}
