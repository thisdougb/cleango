package main

import (
	"github.com/thisdougb/cleango/api/handlers"
	"github.com/thisdougb/cleango/config"
	"github.com/thisdougb/cleango/pkg/datastore/redis"
	"github.com/thisdougb/cleango/pkg/usecase/enablething"
	"log"
	"net/http"
	"os"
)

func main() {

	ds := redis.NewRedisDatastore(config.DB_HOST, config.DB_PORT)

	result := ds.Connect()
	if !result {
		log.Println("Datasore connection failed, exiting...")
		os.Exit(1)
	}
	defer ds.Disconnect()

	env := &handlers.Env{EnableThingService: enablething.NewService(ds)}

	http.HandleFunc("/thing/enable/", env.EnableThing)

	log.Println("webserver.Start(): listening on port", config.API_PORT)
	log.Fatal(http.ListenAndServe(":"+config.API_PORT, nil))
}
