package models

import (
	"time"

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
	CreatedAt   time.Time `json:"createdat" structs:"createdat" bson:"createdat" db:"createdat"`
	UpdatedBy   string `json:"updatedby" structs:"updatedby" bson:"updatedby" db:"updatedby"`
	UpdatedAt   time.Time `json:"updatedat" structs:"updatedat" bson:"updatedat" db:"updatedat"`
}

// Map function returns map values.
func (p *Product) Map() map[string]interface{} {
	return structs.Map(p)
}

// Names function returns field names.
func (p *Product) Names() []string {
	fields := structs.Fields(p)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}
