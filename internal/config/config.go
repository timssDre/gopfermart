package config

import (
	"flag"
	"github.com/caarlos0/env/v10"
	"os"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DataBaseURL string `env:"db"`
	ASA         string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	LogLevel    string `env:"FLAG_LOG_LEVEL"`
}

func InitConfig() *Config {
	dataBaseURL := "host=localhost port=5432 user=postgres password=nbvpass dbname=postgres sslmode=disable"
	config := &Config{
		RunAddress:  "localhost:8080",
		DataBaseURL: dataBaseURL,
		ASA:         "",
		LogLevel:    "info",
	}

	flag.StringVar(&config.RunAddress, "a", config.RunAddress, "address and port to run api")
	flag.StringVar(&config.DataBaseURL, "d", config.DataBaseURL, "address to base store in-memory")
	flag.StringVar(&config.ASA, "r", config.ASA, "nil")
	flag.StringVar(&config.LogLevel, "c", config.LogLevel, "log level")
	flag.Parse()

	if config.DataBaseURL == dataBaseURL {
		dbURI := os.Getenv("DATABASE_URI")
		if dbURI != "" {
			config.DataBaseURL = dbURI
		}

	}
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
