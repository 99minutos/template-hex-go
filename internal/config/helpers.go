package config

import "os"

func IsLogLevel(expected string) bool {
	return os.Getenv("LOG_LEVEL") == expected
}
