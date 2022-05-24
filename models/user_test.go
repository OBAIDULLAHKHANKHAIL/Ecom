package models

import (
	"reflect"
	"testing"
)

func TestUser_Map(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		User     string
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success-student",
			fields: fields{
				ID:       "1234",
				Name:     "obaid",
				User:     "Admin",
				Password: "54321",
			},
			want: map[string]interface{}{
				"id":       "1234",
				"name":     "obaid",
				"user":     "Admin",
				"password": "54321",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Password: tt.fields.Password,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_Names(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		User     string
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				ID:       "1234",
				Name:     "obaid",
				User:     "Admin",
				Password: "54321",
			},
			want: []string{"id", "name", "user", "password"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Password: tt.fields.Password,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
