package main

import (
	"fmt"
	"net/http"
	"test/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello anjay mabar!")
	})
	e.Start(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}