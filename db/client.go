package db

import (
	"log"
	"obaid/models"
)

// DataStore is an interface for query ops.
type DataStore interface {
	AddUser(user *models.User) (string, error)
	UpdateUser(id string, user *models.User) error
	GetUserByID(id string) (*models.User, error)
	RemoveUserByID(id string) error
	AddProduct(product *models.Product) (string, error)
	GetProductByID(id string) (*models.Product, error)
	ListProduct(filter map[string]interface{}, lim int64, off int64) ([]*models.Product, error)
	UpdateProduct(id string, user *models.Product) error
	RemoveProductByID(id string) error
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Reister saves data store into a data store factory.
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
}

