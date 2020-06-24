package dbs

import (
	"go-repository-pattern/config"

	db "gitlab.com/quangdangfit/gocommon/database"
	"gitlab.com/quangdangfit/gocommon/database/mongo"
	"gopkg.in/mgo.v2"
)

var Database mongo.MongoDB

const (
	CollectionProduct = "product"
)

func init() {
	dbConfig := db.DBConfig{
		Hosts:        config.Config.Database.Host,
		AuthDatabase: "admin",
		AuthUserName: config.Config.Database.Username,
		AuthPassword: config.Config.Database.Password,
		Database:     config.Config.Database.Name,
		Env:          config.Config.Database.Env,
		Replica:      config.Config.Database.Replica,
	}

	Database = mongo.New(dbConfig)

	ensureIndex()
}

func ensureIndex() {
	// Index for Product
	indexCode := mgo.Index{
		Key:        []string{"code"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
		Name:       "product_code_index",
	}
	Database.DropIndex(CollectionProduct, indexCode.Name)
	Database.EnsureIndex(CollectionProduct, indexCode)
}
