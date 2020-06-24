package repositories

import (
	"go-repository-pattern/dbs"
	"go-repository-pattern/models"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository interface {
	GetProducts(query bson.M) (*[]models.Product, error)
	GetProduct(query bson.M) (*models.Product, error)
	CreateProduct(product *models.Product) error
}

type productRepository struct {
	repository
}

func NewProductRepository() ProductRepository {
	productRepo := productRepository{}
	productRepo.CollectionName = dbs.CollectionProduct

	return &productRepo
}

func (p *productRepository) GetProducts(query bson.M) (*[]models.Product, error) {
	var product []models.Product
	err := dbs.Database.FindMany(dbs.CollectionProduct, query, "", &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) GetProduct(query bson.M) (*models.Product, error) {
	var product models.Product
	err := dbs.Database.FindOne(dbs.CollectionProduct, query, "", &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) CreateProduct(product *models.Product) error {
	product.BeforeCreate()

	err := dbs.Database.InsertOne(dbs.CollectionProduct, product)
	if err != nil {
		return err
	}
	return nil
}
