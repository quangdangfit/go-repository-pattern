package main

import (
	"context"
	"go-repository-pattern/routers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"transport/lib/utils/logger"
)

//func setupMidlleware()

func main() {
	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano}\t${method}\t${uri}\t${status}\n",
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if code == http.StatusUnauthorized {
			_ = c.JSON(http.StatusUnauthorized, nil)
		}
		e.DefaultHTTPErrorHandler(err, c)
	}

	routers.API(e)

	// Start server
	go func() {
		port := "8000"
		logger.Info("Starting at port: " + port)
		err := e.Start(":" + port)
		if err != nil {
			logger.Error(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
