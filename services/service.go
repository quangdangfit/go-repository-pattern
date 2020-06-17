package services

import "github.com/labstack/echo/v4"

type Service interface {
	List(c echo.Context) (err error)
	Retrieve(c echo.Context) (err error)
	Create(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
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