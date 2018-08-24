package env

import (
	"fmt"
	"os"
)

var (
	LICENSE_DB_PORT = "3306"
	LICENSE_DB_NAME = "license"
	LICENSE_DB_USER = "root"
)

func GetDBHost() (string, error) {
	return MustGet("LICENSE_DB_HOST")
}

func GetDBPort() string {
	return Get("LICENSE_DB_PORT", LICENSE_DB_PORT)
}

func GetDBName() string {
	return Get("LICENSE_DB_NAME", LICENSE_DB_NAME)
}

func GetDBUser() string {
	return Get("LICENSE_DB_USER", LICENSE_DB_USER)
}

func GetDBPass() (string, error) {
	return MustGet("LICENSE_DB_PASS")
}

// Get returns a env value by key.
// If the key does not exist, the default value will be returned.
func Get(key string, defVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defVal
}

// MustGet returns a value by key.
// If the key does not exist, it will return an error.
func MustGet(key string) (string, error) {
	if val := os.Getenv(key); val != "" {
		return val, nil
	}
	return "", fmt.Errorf("no env variable with %s", key)
}
