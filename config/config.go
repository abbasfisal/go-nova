package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port int
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

type RedisConfig struct {
	Host string
	Port int
}

func NewConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		cfg = &Config{
			App: AppConfig{
				Name: getEnv("APP_NAME", "GoApp"),
				Env:  getEnv("APP_ENV", "development"),
				Port: getEnvAsInt("APP_PORT", 8080),
			},
			Database: DatabaseConfig{
				Host: getEnv("DB_HOST", "localhost"),
				Port: getEnvAsInt("DB_PORT", 5432),
				User: getEnv("DB_USER", "root"),
				Pass: getEnv("DB_PASS", ""),
				Name: getEnv("DB_NAME", "appdb"),
			},
			Redis: RedisConfig{
				Host: getEnv("REDIS_HOST", "localhost"),
				Port: getEnvAsInt("REDIS_PORT", 6379),
			},
		}
	})
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
