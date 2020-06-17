package routers

import (
	"github.com/labstack/echo/v4"
	"go-repository-pattern/services"
)

func API(e *echo.Echo) {
	route := e.Group("/api/v1/")

	service := services.NewService()
	route.GET("products", service.List)
}
