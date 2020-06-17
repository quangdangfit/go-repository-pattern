package repositories

import (
	"go-repository-pattern/dbs"
	"go-repository-pattern/models"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository interface {
	Repository

	GetProductByCode(code string) (*models.Product, error)
}

type productRepository struct {
	repository
}

func NewProductRepository() ProductRepository {
	productRepo := productRepository{}
	productRepo.CollectionName = dbs.CollectionProduct

	return &productRepo
}

func (p *productRepository) GetProductByCode(code string) (*models.Product, error) {
	var product models.Product
	query := bson.M{"code": code}
	err := dbs.Database.FindOne(dbs.CollectionProduct, query, "", &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
