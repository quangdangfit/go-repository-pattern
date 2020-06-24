package services

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-repository-pattern/repositories"
	"gopkg.in/mgo.v2/bson"
)

type Service interface {
	List(c echo.Context) (err error)
	Retrieve(c echo.Context) (err error)
	Create(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
	QueryParamsToBson(c echo.Context) (*bson.M, error)
}

type service struct {
	repo repositories.Repository
}

func (s *service) List(c echo.Context) (err error) {
	return nil
}

func (s *service) Retrieve(c echo.Context) (err error) {
	return nil
}

func (s *service) Create(c echo.Context) (err error) {
	return nil
}

func (s *service) Update(c echo.Context) (err error) {
	return nil
}

func (s *service) Delete(c echo.Context) (err error) {
	return nil
}

func (s *service) QueryParamsToBson(c echo.Context) (*bson.M, error) {
	var result bson.M
	payload := map[string]interface{}{}
	queryParams := c.QueryParams()

	for key, _ := range queryParams {
		payload[key] = queryParams.Get(key)
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &result)
	return &result, nil
}
