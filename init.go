package config

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	env := os.Getenv("THOR_ENV")
	if "" == env {
		env = "development"
	}

	_ = godotenv.Load(".env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load()
}