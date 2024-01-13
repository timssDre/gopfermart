package main

import (
	"mBoxMini/internal/app"
	"mBoxMini/internal/config"
)

func main() {
	addrConfig := config.InitConfig()
	appInstance := app.NewApp(addrConfig)
	appInstance.Start()
	appInstance.Stop()
}
