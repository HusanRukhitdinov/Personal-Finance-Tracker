package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT           string
	USER_SERVICE        string
	DB_HOST             string
	DB_PORT             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	SIGNING_KEY         string
	REFRESH_SIGNING_KEY string
	EMAIL               string
	PASSWORD            string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", ":8075"))
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":8081"))
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "3333"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "tracker_auth"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))
	config.REFRESH_SIGNING_KEY = cast.ToString(Coalesce("REFRESH_SIGNING_KEY", "secret"))
	config.EMAIL = cast.ToString(Coalesce("EMAIL", "hrukhitdinov@gmail.com"))
	config.PASSWORD = cast.ToString(Coalesce("PASSWORD", "htyy mkpy wsrk brgg"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
