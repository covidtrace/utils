package env

import (
	"fmt"
	"os"
)

// Get returns an environment variable
func Get(key string) string {
	return os.Getenv(key)
}

// GetDefault returns an environment variable, of a default value if it is not set
func GetDefault(key, def string) string {
	value := Get(key)
	if value == "" {
		return def
	}
	return value
}

// MustGet returns an environment variable, or panics if it is not set
func MustGet(key string) string {
	value := Get(key)
	if value == "" {
		panic(fmt.Errorf("%s is a required env var", key))
	}
	return value
}
