package services

import (
	"github.com/labstack/echo/v4"
	"go-repository-pattern/repositories"
	"net/http"
)

type ProductService interface {
	Service
}

type productService struct {
	service
	repo repositories.ProductRepository
}

func NewProductService() ProductService {
	productSer := productService{}
	productSer.repo = repositories.NewProductRepository()

	return &productSer
}

func (p *productService) List(c echo.Context) (err error) {
	result, err := p.repo.GetProducts()
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, &result)
}

func (p *productService) Retrieve(c echo.Context) (err error) {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := p.repo.GetProductByCode(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, &result)
}
