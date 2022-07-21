package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct{}

// When running mulitple apps in an env, prefix the env vars to give scope separation
const (
	envVarPrefix = "MYAPP_"
)

var defaultValues = map[string]interface{}{
	"API_PORT":         "8080",      // api listens on this port
	"REDIS_HOST":       "localhost", // redis host name
	"REDIS_PORT":       "6379",      // redis port
	"REDIS_USERNAME":   "",          // redis host name
	"REDIS_PASSWORD":   "",          // redis host name
	"REDIS_TLS":        false,       // enables TLS connection to redis
	"REDIS_KEY_PREFIX": "myapp:",    // used to give scope to keys within the redis db
}

// Public methods here.
// Use typed methods so we avoid type assertions at point of use.
func (c *Config) ValueAsStr(key string) string {

	defaultValue := defaultValues[key].(string)
	return c.getEnvVar(key, defaultValue).(string)
}

func (c *Config) ValueAsInt(key string) int {

	defaultValue := defaultValues[key].(int)
	return c.getEnvVar(key, defaultValue).(int)
}

func (c *Config) ValueAsBool(key string) bool {

	defaultValue := defaultValues[key].(bool)
	return c.getEnvVar(key, defaultValue).(bool)
}

// Private methods here
func (c *Config) getEnvVar(key string, fallback interface{}) interface{} {

	fullEnvVarName := fmt.Sprintf("%s%s", envVarPrefix, key)
	value, exists := os.LookupEnv(fullEnvVarName)
	if !exists {
		return fallback
	}

	switch fallback.(type) {
	case string:
		return value
	case bool:
		valueAsBool, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return valueAsBool
	case int:
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return valueAsInt
	}
	return fallback
}
