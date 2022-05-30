package db

import (
	"log"

	"github.com/OBAIDULLAHKHANKHAIL/Ecom/models"
)

// DataStore is an interface for query ops.
type DataStore interface {
	AddStudent(product *models.Product) (string, error)
	UpdateStudent(id string, student *models.Product) error
	GetStudentByID(id string) (*models.Product, error)
	ListStudent(filter map[string]interface{}, lim int64, off int64) ([]*models.Product, error)
	RemoveStudentByID(id string) error
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory.
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