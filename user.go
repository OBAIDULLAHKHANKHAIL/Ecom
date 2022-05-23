package models

import (
	"github.com/fatih/structs"
)

// User structs.
type User struct {
	ID       string `json:"id" structs:"id" bson:"_id" db:"id"`
	Name     string `json:"name" structs:"name" bson:"name" db:"name"`
	User     string `json:"user" structs:"user" bson:"user" db:"user"`
	Password string `json:"password" structs:"password" bson:"password" db:"password"`
}

// Map function returns map values.
func (u *User) Map() map[string]interface{} {
	return structs.Map(u)
}
