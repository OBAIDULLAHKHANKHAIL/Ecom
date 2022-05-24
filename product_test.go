package models

import (
	"reflect"
	"testing"
)

func TestProduct_Map(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Price       int
		Description string
		Password    string
		CreatedBy   string
		CreatedAt   string
		UpdatedBy   string
		UpdatedAt   string
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
				CreatedAt:   "24-may-22",
				UpdatedBy:   "Ali",
				UpdatedAt:   "25-may-22",
			},
			want: map[string]interface{}{
				"id":          "1234",
				"name":        "T-Shirt",
				"price":       2900,
				"description": "Full Seleves T-shirt",
				"password":    "9876",
				"createdby":   "Obaid",
				"createdat":   "24-may-22",
				"updatedby":   "Ali",
				"updatedat":   "25-may-22",
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
		CreatedAt   string
		UpdatedBy   string
		UpdatedAt   string
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
				CreatedAt:   "24-may-22",
				UpdatedBy:   "Ali",
				UpdatedAt:   "25-may-22",
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
