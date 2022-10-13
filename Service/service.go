package service

import (
	"obaid/models"
)

// AddUser adds or update user into database.
func (s *Service) AddUser(user *models.User) (string, error) {
	return s.db.AddUser(user)
}


// UpdateUser updates User record in database.
func (s *Service) UpdateUser(id string, user *models.User) error {
	return s.db.UpdateUser(id, user)
}

// GetUser retrieve user from database.
func (s *Service) GetUser(id string) (*models.User, error) {
	return s.db.GetUserByID(id)
}

// DeleteUser deletes user from database.
func (s *Service) DeleteUser(id string) error {
	return s.db.RemoveUserByID(id)
}

// AddProduct adds or update product into database.
func (s *Service) AddProduct(product *models.Product) (string, error) {
	return s.db.AddProduct(product)
}

// UpdateProduct updates product record in database.
func (s *Service) UpdateProduct(id string, product *models.Product) error {
	return s.db.UpdateProduct(id, product)
}

// GetProduct retrieve product from database.
func (s *Service) GetProduct(id string) (*models.Product, error) {
	return s.db.GetProductByID(id)
}

// DeleteProduct deletes product from database.
func (s *Service) DeleteProduct(id string) error {
	return s.db.RemoveProductByID(id)
}

// ListStudents retrieve all the students from database.
func (s *Service) ListProduct(filter map[string]interface{}, lim int64, off int64) ([]*models.Product, error) {
	return s.db.ListProduct(filter, lim, off)
}