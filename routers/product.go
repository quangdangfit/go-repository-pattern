package routers

import (
	"github.com/labstack/echo/v4"
	"go-repository-pattern/services"
)

func API(e *echo.Echo) {
	route := e.Group("/api/v1/")

	product := services.NewProductService()
	route.GET("products", product.List)
	route.GET("products/:code", product.Retrieve)
}
