package models

import (
	"github.com/fatih/structs"
)

// Product struct.
type Product struct {
	ID          string `json:"id" structs:"id" bson:"_id" db:"id"`
	Name        string `json:"name" structs:"name" bson:"name" db:"name"`
	Price       int    `json:"price" structs:"price" bson:"price" db:"price"`
	Description string `json:"description" structs:"description" bson:"description" db:"description"`
	Password    string `json:"password" structs:"password" bson:"password" db:"password"`
	CreatedBy   string `json:"createdby" structs:"createdby" bson:"createdby" db:"createdby"`
	CreatedAt   string `json:"createdat" structs:"createdat" bson:"createdat" db:"createdat"`
	UpdatedBy   string `json:"updatedby" structs:"updatedby" bson:"updatedby" db:"updatedby"`
	UpdatedAt   string `json:"updatedat" structs:"updatedat" bson:"updatedat" db:"updatedat"`
}

// Map function returns map values.
func (p *Product) Map() map[string]interface{} {
	return structs.Map(p)
}
