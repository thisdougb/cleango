package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Datastore struct {
	host     string
	port     string
	username string
	password string
	database string
	tls      bool
	ctx      context.Context
	client   *redis.Client
}

func NewRedisDatastore(host string, port string, username string, password string, tls bool) *Datastore {
	return &Datastore{host, port, username, password, "0", tls, context.Background(), nil}
}
