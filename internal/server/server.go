package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/von-salumbides/go-sls-crud-api/utils/logger"
	"go.uber.org/zap"
)

type ServerConfig struct {
	HttpAddress string
	IsCors      bool
	IsDebug     bool
}

func NewServerConfig(httpAddress string, isCors bool, isDebug bool) *ServerConfig {
	return &ServerConfig{
		HttpAddress: httpAddress,
		IsCors:      isCors,
		IsDebug:     isDebug,
	}
}

type Server struct {
	Echo        *echo.Echo
	HttpAddress string
}

func NewServer(config *ServerConfig) (*Server, error) {
	e := echo.New()
	e.Use(middleware.CSRF())
	// Log HTTP requests
	e.Use(middleware.Logger())
	// Recover
	e.Use(middleware.Recover())
	// Cors
	if config.IsCors {
		logger.INFO("CORS has been enabled", nil)
		e.Use(middleware.CORS())
	}

	return &Server{
		Echo:        e,
		HttpAddress: config.HttpAddress,
	}, nil
}

func (server *Server) Start() {
	zap.L().Info("Starting HTTP server",
		zap.Any("address:", server.HttpAddress),
	)

	err := server.Echo.Start(server.HttpAddress)
	if err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("Failed to start server",
			zap.String("address", server.HttpAddress),
			zap.Error(err))
	}
}

func (server *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	zap.L().Info("Shutting down HTTP server",
		zap.String("address", server.HttpAddress))
	if err := server.Echo.Shutdown(ctx); err != nil {
		zap.L().Error("Failed to shut down HTTP server gracefully", zap.Error(err))
		return err
	}
	return nil
}
