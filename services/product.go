package services

import (
	"go-repository-pattern/models"
	"go-repository-pattern/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
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
	req := models.ProductRequest{}
	query, _ := p.ToBson(req)

	result, err := p.repo.GetProducts(query)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, &result)
}

func (p *productService) Retrieve(c echo.Context) (err error) {
	req := models.ProductRequest{}
	query, _ := p.ToBson(req)

	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := p.repo.GetProduct(query)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, &result)
}
