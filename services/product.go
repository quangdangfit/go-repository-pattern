package services

import (
	"go-repository-pattern/models"
	"go-repository-pattern/repositories"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
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
	query, _ := p.QueryParamsToBson(c)

	result, err := p.repo.GetProducts(*query)
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

	query := bson.M{"code": code}
	result, err := p.repo.GetProduct(query)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, &result)
}

func (p *productService) Create(c echo.Context) (err error) {
	var req models.ProductRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("SCHEDULE: Bad request: ", err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		logger.Error("SCHEDULE: Bad request: ", err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	product := models.Product{}
	copier.Copy(&product, &req)

	err = p.repo.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, &product)
}
