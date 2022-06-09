package mongo

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"obaid/db"
	"obaid/models"
)

func Test_mongoStd_AddUser(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add user in db",
			args: args{user: &models.User{
				Name:     "bilal",
				User:     "dev12",
				Password: "123111s",
			}},
			wantErr: false,
		},
		{
			name: "fail - add invalid user in db",
			args: args{user: &models.User{
				Name:     "bilal",
				User:     "dev12",
				Password: "123111s",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMongoClient(db.Option{})
			if err != nil {
				panic(err)
			}
			_, err = m.AddUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mongoStd_UpdateUser(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewMongoClient(db.Option{})
	user := &models.User{
		Name:     "bilal",
		User:     "dev12",
		Password: "123111s",
	}
	_, _ = m.AddUser(user)
	type args struct {
		id   string
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
					Name:     "bilal",
					User:     "dev12",
					Password: "123111s",
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

func Test_mongoStd_GetUserByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewMongoClient(db.Option{})
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

func Test_mongoStd_RemoveUserByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewMongoClient(db.Option{})
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
			name:    "success - delete user by ID",
			args:    args{id: user.ID},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.RemoveUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mongoStd_AddProduct(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	type args struct {
		product *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add product in db",
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
		{
			name: "fail - add invalid product in db",
			args: args{product: &models.Product{
				Name:        "Tea free",
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
			m, _ := NewMongoClient(db.Option{})
			_, err := m.AddProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mongoStd_UpdateProduct(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewMongoClient(db.Option{})
	product := &models.Product{
		Name:        "Tea",
		Price:       499,
		Description: "Pakistan #1 Tea",
		Password:    "1122",
		CreatedBy:   "obaid",
		CreatedAt:   time.Now(),
		UpdatedBy:   "Umair",
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
					Name:        "Tea",
					Price:       499,
					Description: "Pakistan #1 Tea",
					Password:    "1122",
					CreatedBy:   "obaid",
					CreatedAt:   time.Now(),
					UpdatedBy:   "Umair",
					UpdatedAt:   time.Now(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.UpdateProduct(tt.args.id, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoStd_GetProductByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewMongoClient(db.Option{})
	product := &models.Product{
		Name:        "Black Tea",
		Price:       499,
		Description: "Turkey #1 Tea",
		Password:    "1133",
		CreatedBy:   "Ali",
		CreatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
		UpdatedBy:   "Usman",
		UpdatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
	}
	_, _ = m.AddProduct(product)
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
			name:    "success - get product by ID",
			args:    args{id: product.ID},
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

func Test_mongoStd_RemoveProductByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewMongoClient(db.Option{})
	product := &models.Product{
		Name:        "Black Tea",
		Price:       499,
		Description: "Turkey #1 Tea",
		Password:    "1133",
		CreatedBy:   "Ali",
		CreatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
		UpdatedBy:   "Usman",
		UpdatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
	}
	_, _ = m.AddProduct(product)
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
			name:    "success - delete product by ID",
			args:    args{id: product.ID},
			want:    product,
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

func Test_mongoStd_ListProduct(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	m, _ := NewMongoClient(db.Option{})
	product1 := &models.Product{
		Name:        "Black Tea",
		Price:       499,
		Description: "Turkey #1 Tea",
		Password:    "1133",
		CreatedBy:   "Ali",
		CreatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
		UpdatedBy:   "Usman",
		UpdatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
	}
	product2 := &models.Product{
		Name:        "Green Tea",
		Price:       499,
		Description: "Pakistan #1 Tea",
		Password:    "1144",
		CreatedBy:   "zahid",
		CreatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
		UpdatedBy:   "saqib",
		UpdatedAt:   time.Date(2022, 04, 02, 12, 12, 2, 0, time.UTC),
	}
	_, _ = m.AddProduct(product1)
	_, _ = m.AddProduct(product2)
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
			name:    "success- list all product with name",
			args:    args{filter: map[string]interface{}{"name": "Black Tea"}, lim: 1, off: 0},
			want:    []*models.Product{product1},
			wantErr: false,
		},
		{
			name:    "fail- list all product with name",
			args:    args{filter: map[string]interface{}{"name": "Greens Tea"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all product with name createdby",
			args:    args{filter: map[string]interface{}{"name": "Black Tea", "createdby": "Ali"}, lim: 1, off: 0},
			want:    []*models.Product{product1},
			wantErr: false,
		},
		{
			name:    "fail- list all students with name createdby",
			args:    args{filter: map[string]interface{}{"name": "Green Tea", "level": "zahid"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all students with updatedby ",
			args:    args{filter: map[string]interface{}{"updatedby": "saqib"}, lim: 1, off: 0},
			want:    []*models.Product{product2},
			wantErr: false,
		},
		{
			name:    "fail- list all students with updatedby",
			args:    args{filter: map[string]interface{}{"updatedby": "zahid"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all students with name price",
			args:    args{filter: map[string]interface{}{"name": "Green Tea", "price": 499}, lim: 1, off: 0},
			want:    []*models.Product{product2},
			wantErr: false,
		},
		{
			name:    "fail- list all students with name price",
			args:    args{filter: map[string]interface{}{"name": "Black Tea", "price": 500}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
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
