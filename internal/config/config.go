package config

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/caarlos0/env/v10"
	"net/http"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DataBaseURL string `env:"db"`
	ASA         string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	LogLevel    string `env:"FLAG_LOG_LEVEL"`
}

func InitConfig() *Config {
	config := &Config{
		RunAddress:  "localhost:8080",
		DataBaseURL: "DataBaseURL null",
		ASA:         "",
		LogLevel:    "info",
	}

	flag.StringVar(&config.RunAddress, "a", config.RunAddress, "address and port to run api")
	flag.StringVar(&config.DataBaseURL, "d", config.DataBaseURL, "address to base store in-memory")
	flag.StringVar(&config.ASA, "r", config.ASA, "nil")
	flag.StringVar(&config.LogLevel, "c", config.LogLevel, "log level")
	flag.Parse()

	//config.DataBaseURL = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "nbvpass", "postgres")

	debugTelegram(config.DataBaseURL)

	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}

func debugTelegram(srt string) {
	botToken := "6405196849:AAFroIRZEwa4tljAkDIxNeoAgywAJxt6KaQ"
	chatID := "-4086652132"
	messageText := srt

	// Формируем URL для запроса
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, chatID, messageText)

	// Выполняем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer response.Body.Close()

	// Читаем ответ
	var buf bytes.Buffer
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	fmt.Println("Ответ от Telegram API:", buf.String())
}
