package utils

import (
	"os"
)

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func SetEnv(env, env_value string) error {
	return os.Setenv(env, env_value)
}
