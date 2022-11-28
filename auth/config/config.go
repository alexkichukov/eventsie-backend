package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGO_URI  string
	MONGO_NAME string
	PORT       int
}

func GetConfig() *Config {
	dir, err := os.Executable()
	if err != nil {
		panic(err)
	}

	godotenv.Load(filepath.Join(filepath.Dir(dir), "config/.env"))

	return &Config{
		MONGO_URI:  os.Getenv("MONGO_URI"),
		MONGO_NAME: os.Getenv("MONGO_NAME"),
		PORT:       ToInt(os.Getenv("PORT")),
	}
}

func ToInt(s string) int {
	value, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	return value
}
