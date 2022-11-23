package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	API_PORT            int
	EVENTS_SERVICE_PORT int
	AUTH_SERVICE_PORT   int
}

func GetConfig() *Config {
	dir, err := os.Executable()
	if err != nil {
		panic(err)
	}

	godotenv.Load(filepath.Join(filepath.Dir(dir), "config/.env"))

	return &Config{
		API_PORT:            ToInt(os.Getenv("API_PORT")),
		EVENTS_SERVICE_PORT: ToInt(os.Getenv("EVENTS_SERVICE_PORT")),
		AUTH_SERVICE_PORT:   ToInt(os.Getenv("AUTH_SERVICE_PORT")),
	}
}

func ToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return value
}
