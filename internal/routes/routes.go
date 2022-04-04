package routes

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/von-salumbides/go-sls-crud-api/internal/server"
)

func Routes(server *server.Server) {
	r := server.Echo
	r.GET("/health", Health)
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"healthy": "ok",
		"runtimeData": map[string]int{
			"cpu":       runtime.NumCPU(),
			"goRoutine": runtime.NumGoroutine(),
		},
	})
}
