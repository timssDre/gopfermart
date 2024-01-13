package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"mBoxMini/internal/logger"
	"mBoxMini/internal/services"
	"mBoxMini/repository"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type RestAPI struct {
	BoxService *services.BoxService
}

func StartRestAPI(serverAddr string, LogLevel string, db *repository.StoreDB) error {
	if err := logger.Initialize(LogLevel); err != nil {
		return err
	}
	logger.Log.Info("Running server", zap.String("address", serverAddr))
	BoxService := services.NewBoxService(db)

	api := &RestAPI{
		BoxService,
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		gin.Recovery(),
	)
	api.setRoutes(r)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		err := r.Run(serverAddr)
		if err != nil {
			fmt.Println("failed to start the browser")
		}
	}()
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при остановке сервера: %v\n", err)
	}

	return nil
}
