package config

import (
	"flag"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DataBaseURl string `env:"db"`
	ASA         string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func InitConfig() *Config {
	config := &Config{
		RunAddress:  "localhost:8080",
		DataBaseURl: "",
		ASA:         "",
	}

	flag.StringVar(&config.RunAddress, "a", config.RunAddress, "address and port to run api")
	flag.StringVar(&config.DataBaseURl, "d", config.DataBaseURl, "address to base store in-memory")
	flag.StringVar(&config.ASA, "r", config.ASA, "nil")
	flag.Parse()

	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
