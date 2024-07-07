package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPAssword string
	DBAddress  string
	DBNAme     string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPAssword: getEnv("DB_PASSWORD", ""),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_HOST", "3306")),
		DBNAme:     getEnv("DB_NAME", "ketoo"),
	}
}

func getEnv(Key, fallback string) string {
	if value, ok := os.LookupEnv(Key); ok {
		return value
	}
	return fallback
}
