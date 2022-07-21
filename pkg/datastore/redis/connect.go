package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// Connect to set Redis connection string
func (d *Datastore) Connect() bool {

	connection := fmt.Sprintf("%s:%s", d.host, d.port)

	i, err := strconv.Atoi(d.database)
	if err != nil {
		return false
	}

	d.client = redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       i,
	})

	_, err = d.client.Ping(d.ctx).Result()
	if err != nil {
		fmt.Println("redis connect error: ", err)
		return false
	}

	return true
}

func (d *Datastore) Disconnect() {
	d.client.Close()
}
