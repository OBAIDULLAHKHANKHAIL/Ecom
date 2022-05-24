package models

import (
	"reflect"
	"testing"
	"time"
)

func TestProduct_Map(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Price       int
		Description string
		Password    string
		CreatedBy   string
		CreatedAt   time.Time
		UpdatedBy   string
		UpdatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success-student",
			fields: fields{
				ID:          "1234",
				Name:        "T-Shirt",
				Price:       2900,
				Description: "Full Seleves T-shirt",
				Password:    "9876",
				CreatedBy:   "Obaid",
				CreatedAt:   time.Now(),
				UpdatedBy:   "Ali",
				UpdatedAt:   time.Now(),
			},
			want: map[string]interface{}{
				"id":          "1234",
				"name":        "T-Shirt",
				"price":       2900,
				"description": "Full Seleves T-shirt",
				"password":    "9876",
				"createdby":   "Obaid",
				"createdat":   time.Now(),
				"updatedby":   "Ali",
				"updatedat":   time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Product{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Price:       tt.fields.Price,
				Description: tt.fields.Description,
				Password:    tt.fields.Password,
				CreatedBy:   tt.fields.CreatedBy,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedBy:   tt.fields.UpdatedBy,
				UpdatedAt:   tt.fields.CreatedAt,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_Names(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Price       int
		Description string
		Password    string
		CreatedBy   string
		CreatedAt   time.Time
		UpdatedBy   string
		UpdatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				ID:          "1234",
				Name:        "T-Shirt",
				Price:       2900,
				Description: "Full Seleves T-shirt",
				Password:    "9876",
				CreatedBy:   "Obaid",
				CreatedAt:   time.Now(),
				UpdatedBy:   "Ali",
				UpdatedAt:   time.Now(),
			},
			want: []string{"id", "name", "price", "description", "password", "createdby", "createdat", "updatedby", "updatedat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Product{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Price:       tt.fields.Price,
				Description: tt.fields.Description,
				Password:    tt.fields.Password,
				CreatedBy:   tt.fields.CreatedBy,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedBy:   tt.fields.UpdatedBy,
				UpdatedAt:   tt.fields.CreatedAt,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
