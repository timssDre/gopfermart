package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DataBaseURl string `env:"db"`
	ASA         string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	LogLevel    string `env:"FLAG_LOG_LEVEL"`
}

func InitConfig() *Config {
	config := &Config{
		RunAddress:  "localhost:8080",
		DataBaseURl: "",
		ASA:         "",
		LogLevel:    "info",
	}

	flag.StringVar(&config.RunAddress, "a", config.RunAddress, "address and port to run api")
	flag.StringVar(&config.DataBaseURl, "d", config.DataBaseURl, "address to base store in-memory")
	flag.StringVar(&config.ASA, "r", config.ASA, "nil")
	flag.StringVar(&config.LogLevel, "c", config.LogLevel, "log level")
	flag.Parse()

	config.DataBaseURl = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "nbvpass", "postgres")

	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
