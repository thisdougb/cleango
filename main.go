package main

import (
	"log"
	"net/http"
	"time"

	"github.com/thisdougb/cleango/api/handlers"
	"github.com/thisdougb/cleango/config"
	"github.com/thisdougb/cleango/pkg/datastore/redis"
	"github.com/thisdougb/cleango/pkg/usecase/enablething"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	var cfg *config.Config // dynamic config settings

	ds := redis.NewRedisDatastore(
		cfg.ValueAsStr("REDIS_HOST"),
		cfg.ValueAsStr("REDIS_PORT"),
		cfg.ValueAsStr("REDIS_USERNAME"),
		cfg.ValueAsStr("REDIS_PASSWORD"),
		cfg.ValueAsBool("REDIS_TLS"))

	// To be container friendly loop until we get a datastore connection. Otherwise the container
	// goes into crash-loop and it's harder to troubleshoot.
	for {
		log.Printf("Datastore connecting, host: '%s:%s', username: %s\n",
			cfg.ValueAsStr("REDIS_HOST"),
			cfg.ValueAsStr("REDIS_PORT"),
			cfg.ValueAsStr("REDIS_USERNAME"))

		err := ds.Connect()
		if err == nil {
			log.Println("Datastore connected.")
			break
		}
		log.Println("Datastore connect failed:", err.Error())
		time.Sleep(5 * time.Second)
	}
	defer ds.Disconnect()

	env := &handlers.Env{EnableThingService: enablething.NewService(ds)}

	http.HandleFunc("/thing/enable/", env.EnableThing)

	log.Println("webserver.Start(): listening on port", cfg.ValueAsStr("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+cfg.ValueAsStr("API_PORT"), nil))
}
