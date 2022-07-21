// +build dev

package config

import (
	"os"
)

func GetTemplatePath(fileName string) string {

	if _, ok := os.LookupEnv("GOPATH"); ok {
		return os.Getenv("GOPATH") + "/src/github.com/thisdougb/gonelong/api/templates/" + fileName
	}

	cwd, _ := os.Getwd()
	return cwd + "/../templates/" + fileName // for handler unit tests
}
