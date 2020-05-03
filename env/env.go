package env

import (
	"fmt"
	"os"
)

func Get(key string) string {
	return os.Getenv(key)
}

func MustGet(key string) string {
	value := Get(key)
	if value == "" {
		panic(fmt.Errorf("%s is a required env var", key))
	}
	return value
}
