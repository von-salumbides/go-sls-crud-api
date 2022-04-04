package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/von-salumbides/go-sls-crud-api/internal/server"
	"github.com/von-salumbides/go-sls-crud-api/utils/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	serverConfig := server.NewServerConfig(":9000", false, false)
	httpServer, err := server.NewServer(serverConfig)
	if err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("Shutting down the server", zap.Error(err))
		os.Exit(1)
	}
	go httpServer.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	err = httpServer.Shutdown(300 * time.Second)
	if err == nil {
		zap.L().Info("Shutdown Successfully")
	}
}
