package app

import (
	"log"
	"mBoxMini/internal/api"
	"mBoxMini/internal/config"
	"mBoxMini/repository"
)

type App struct {
	ServerAddress string
	config        *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) Start() {
	db, err := repository.InitDatabase(a.config.DataBaseURl)
	if err != nil {
		log.Fatal(err)
	}
	err = api.StartRestAPI(a.config.RunAddress, a.config.LogLevel, db)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Stop() {

}
