package mysql

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"obaid/db"
	"obaid/models"
)

func Test_mysqlStd_AddUser(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success - add user in db",
			args: args{user: &models.User{
				Name:     "Obaid",
				User:     "dev",
				Password: "12314",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMysqlClient(db.Option{})
			if err != nil {
				fmt.Println(err)
				panic("fail")
			}
			id, err := m.AddUser(tt.args.user)
			if err != nil {
				t.Errorf("AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(id)
		})
	}
}

func Test_mysqlStd_UpdateUser(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	m, _ := NewMysqlClient(db.Option{})
	user := &models.User{
		Name:     "Ali",
		User:     "dev2",
		Password: "1231s",
	}
	_, _ = m.AddUser(user)
	type args struct {
		id      string
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - update records",
			args: args{id: user.ID,
				user: &models.User{
					ID: user.ID,
					Name:     "Ali",
					User:     "dev23",
					Password: "1231s",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.UpdateUser(tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlStd_GetUserByID(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	m, _ := NewMysqlClient(db.Option{})
	user := &models.User{
		Name:     "bilal",
		User:     "dev12",
		Password: "123111s",
	}
	_, _ = m.AddUser(user)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "success - get user by ID",
			args:    args{id: user.ID},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if cmp.Equal(got, tt.want) {
				log().Info(got)
			} else {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStd_AddProduct(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	type args struct {
		product *models.Product
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success - add Product in db",
			args: args{product: &models.Product{
				Name:        "Tea",
				Price:       499,
				Description: "Pakistan #1 Tea",
				Password:    "1122",
				CreatedBy:   "obaid",
				CreatedAt:   time.Now(),
				UpdatedBy:   "Umair",
				UpdatedAt:   time.Now(),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMysqlClient(db.Option{})
			if err != nil {
				fmt.Println(err)
				panic("fail")
			}
			id, err := m.AddProduct(tt.args.product)
			if err != nil {
				t.Errorf("AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(id)
		})
	}
}

func Test_mysqlStd_GetProductByID(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	m, _ := NewMysqlClient(db.Option{})
	product := &models.Product{
		Name:        "Black Tea",
		Price:       499,
		Description: "Turkey #1 Tea",
		Password:    "1133",
		CreatedBy:   "Ali",
		CreatedAt:   time.Date(2022,04,02,12,12,2,0,time.UTC),
		UpdatedBy:   "Usman",
		UpdatedAt:   time.Date(2022,04,02,12,12,2,0,time.UTC),
	}
	id, err := m.AddProduct(product)
	if err != nil {
		panic(err)
	}
	product.ID = id
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Product
		wantErr bool
	}{
		{
			name:    "success - get Product by ID",
			args:    args{id: id},
			want:    product,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetProductByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if cmp.Equal(got, tt.want) {
				log().Info(got)
			} else {
				t.Errorf("GetProductByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStd_ListProduct(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	m, _ := NewMysqlClient(db.Option{})
	product := &models.Product{
		Name:        "Green Tea",
		Price:       499,
		Description: "Pakistan #1 Tea",
		Password:    "1144",
		CreatedBy:   "zahid",
		CreatedAt:   time.Date(2022,04,02,12,12,2,0,time.UTC),
		UpdatedBy:   "saqib",
		UpdatedAt:   time.Date(2022,04,02,12,12,2,0,time.UTC),
	}
	Product2 := &models.Product{
		Name:        "Pink Tea",
		Price:       499,
		Description: "US #1 Tea",
		Password:    "911",
		CreatedBy:   "zahid",
		CreatedAt:   time.Date(2022,04,02,12,12,2,0,time.UTC),
		UpdatedBy:   "ladin",
		UpdatedAt:   time.Date(2022,04,02,12,12,2,0,time.UTC),
	}
	_, _ = m.AddProduct(product)
	_, _ = m.AddProduct(Product2)
	type args struct {
		filter map[string]interface{}
		lim    int64
		off    int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Product
		wantErr bool
	}{
		{
			name:    "success- list all Product with name",
			args:    args{filter: map[string]interface{}{"name": "saqib"}, lim: 1, off: 0},
			want:    []*models.Product{product},
			wantErr: false,
		},
		{
			name:    "fail- list all Product with name",
			args:    args{filter: map[string]interface{}{"name": "zahid"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all product with name pass",
			args:    args{filter: map[string]interface{}{"name": "saqib", "password": "114"}, lim: 1, off: 0},
			want:    []*models.Product{Product2},
			wantErr: false,
		},
		{
			name:    "fail- list all user with name pass",
			args:    args{filter: map[string]interface{}{"name": "zahid", "password": "911"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all user with UpdatedBy ",
			args:    args{filter: map[string]interface{}{"UpdatedBy": "saqib"}, lim: 1, off: 0},
			want:    []*models.Product{Product2},
			wantErr: false,
		},
		{
			name:    "fail- list all user with phone",
			args:    args{filter: map[string]interface{}{"UpdatedBy": "zahid"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		// {
		// 	name:    "success- list all user with phone CreatedBy",
		// 	args:    args{filter: map[string]interface{}{"CreatedBy": "911", "phone": "03217895641"}, lim: 1, off: 0},
		// 	want:    []*models.Product{product},
		// 	wantErr: false,
		// },
		// {
		// 	name:    "fail- list all user with name Hameed CreatedBy",
		// 	args:    args{filter: map[string]interface{}{"CreatedBy": "114", "phone": "031458795149"}, lim: 1, off: 0},
		// 	want:    nil,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.ListProduct(tt.args.filter, tt.args.lim, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ListProduct() got = %v, want = %v,diff = %s, error = %v", got, tt.want, diff, err)
			} else {
				log().Info(got)
			}
		})
	}
}
func Test_mysqlStd_UpdateProduct(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")

	m, _ := NewMysqlClient(db.Option{})
	product := &models.Product{
		Name:        "Green Tea",
		Price:       499,
		Description: "Pakistan #1 Tea",
		Password:    "1144",
		CreatedBy:   "zahid",
		CreatedAt:   time.Now(),
		UpdatedBy:   "saqib",
		UpdatedAt:   time.Now(),
	}
	_, _ = m.AddProduct(product)
	type args struct {
		id      string
		product *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - update records",
			args: args{id: product.ID,
				product: &models.Product{
					Name:        "Black Tea",
					Price:       499,
					Description: "Pakistan #1 Tea",
					Password:    "1144",
					CreatedBy:   "zahid",
					CreatedAt:   time.Now(),
					UpdatedBy:   "saqib",
					UpdatedAt:   time.Now(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.UpdateProduct(tt.args.id, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



func Test_mysqlStd_RemoveProductByID(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	m, _ := NewMysqlClient(db.Option{})
	product := &models.Product{
		Name:        "Green Tea",
		Price:       499,
		Description: "Pakistan #1 Tea",
		Password:    "1144",
		CreatedBy:   "zahid",
		CreatedAt:   time.Now(),
		UpdatedBy:   "saqib",
		UpdatedAt:   time.Now(),
	}
	_, _ = m.AddProduct(product)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete Product by ID",
			args:    args{id: product.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.RemoveProductByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

