package config

import "os"

func IsProd() bool {
	return os.Getenv("APP_ENV") == "prod"
}

func GetEnv() string {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "dev")
		return "dev"
	}

	return os.Getenv("APP_ENV")
}
