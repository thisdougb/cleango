package redis

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// Connect to set Redis connection string
func (d *Datastore) Connect() error {

	connection := fmt.Sprintf("%s:%s", d.host, d.port)

	i, err := strconv.Atoi(d.database)
	if err != nil {
		return err
	}

	if d.tls {

		d.client = redis.NewClient(&redis.Options{
			Addr:     connection,
			Username: d.username,
			Password: d.password,
			DB:       i,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
				//Certificates: []tls.Certificate{cert}
			},
		})
	} else {

		d.client = redis.NewClient(&redis.Options{
			Addr:     connection,
			Password: "",
			DB:       i,
		})
	}

	_, err = d.client.Ping(d.ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (d *Datastore) Disconnect() {
	d.client.Close()
}
