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

// Names function returns field names.
func (u *User) Names() []string {
	fields := structs.Fields(u)
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