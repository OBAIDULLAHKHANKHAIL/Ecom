package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"obaid/config"
	"obaid/db"
	"obaid/models"
)

type client struct {
	db *sqlx.DB
}

func init() {
	db.Register("mysql", NewMysqlClient)
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

//NewMysqlClient to connect DB.
func NewMysqlClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	return &client{db: cli}, nil
}

func (m client) AddUser(user *models.User) (string, error) {
	if user.ID != "" {
		return "id is not empty", nil
	}
	user.ID = guuid.NewV4().String()
	names := user.Names()
	if _, err := m.db.NamedExec(fmt.Sprintf(`INSERT INTO user (%s) VALUES(%s)`, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")),
		user); err != nil {
		return "", errors.Wrap(err, "failed to add User")
	}
	return user.ID, nil
}

func (m client) UpdateUser(id string, user *models.User) error {
	names := user.Names()
	if _, err := m.db.NamedExec(fmt.Sprintf(`UPDATE user SET %s WHERE id = '%s'`, strings.Join(mkPlaceHolder(names, "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ","), id), user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (m client) GetUserByID(id string) (*models.User, error) {
	var std models.User
	if err := m.db.Get(&std, fmt.Sprintf(`SELECT * FROM user WHERE id='%s'`, id)); err != nil {
		if sql.ErrNoRows != nil {
			return nil, errors.Wrap(err, "failed to fetch User....not found")
		}
		return nil, err
	}
	fmt.Printf("%v\n", std.Map())
	return &std, nil
}

func (m client) RemoveUserByID(id string) error {
	if _, err := m.db.Exec(fmt.Sprintf(`DELETE FROM user WHERE id= '%s'`, id)); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}

func (m client) AddProduct(product *models.Product) (string, error) {
	if product.ID != "" {
		return "id is not empty", nil
	}
	product.ID = guuid.NewV4().String()
	names := product.Names()
	if _, err := m.db.NamedExec(fmt.Sprintf(`INSERT INTO product (%s) VALUES(%s)`, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")),
		product); err != nil {
		return "", errors.Wrap(err, "failed to add product")
	}
	return product.ID, nil
}

func (m client) GetProductByID(id string) (*models.Product, error) {
	var std models.Product
	if err := m.db.Get(&std, fmt.Sprintf(`SELECT * FROM product WHERE ID='%s'`, id)); err != nil {
		if sql.ErrNoRows != nil {
			return nil, errors.Wrap(err, "failed to fetch Product....not found")
		}
		return nil, err
	}
	fmt.Printf("%v\n", std.Map())
	return &std, nil
}

func (m client) ListProduct(filter map[string]interface{}, lim int64, off int64) ([]*models.Product, error) {
	var stdList []*models.Product
	if err := m.db.Select(&stdList, fmt.Sprintf("SELECT * FROM product %s LIMIT %d,%d", mkFilter(filter), off, lim)); err != nil {
		if sql.ErrNoRows != nil {
			return stdList, errors.Wrap(err, "failed to list product....not found")
		}
		return stdList, err
	}
	log().Info(stdList)
	return stdList, nil
}

func (m client) RemoveProductByID(id string) error {
	if _, err := m.db.Exec(fmt.Sprintf(`DELETE FROM product WHERE id= '%s'`, id)); err != nil {
		return errors.Wrap(err, "failed to delete product")
	}
	return nil
}

func (m client) UpdateProduct(id string, product *models.Product) error {
	names := product.Names()
	if _, err := m.db.NamedExec(fmt.Sprintf(`UPDATE product SET %s WHERE id = '%s'`, strings.Join(mkPlaceHolder(names, "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ","), id), product); err != nil {
		return errors.Wrap(err, "failed to update product")
	}
	return nil
}

// func (m client) GetProductByID(id string) (*models.Product, error) {
// 	var std models.Product
// 	if err := m.db.Get(&std, fmt.Sprintf(`SELECT * FROM product WHERE id='%s'`, id)); err != nil {
// 		if sql.ErrNoRows != nil {
// 			return nil, errors.Wrap(err, "failed to fetch product....not found")
// 		}
// 		return nil, err
// 	}
// 	fmt.Printf("%v\n", std.Map())
// 	return &std, nil
// }

func mkFilter(fil map[string]interface{}) string {
	if len(fil) < 1 {
		return ""
	}

	whereFilter := make([]string, 0)
	for key, val := range fil {
		whereFilter = append(whereFilter, fmt.Sprintf(" %s='%s' ", key, val))
	}

	return fmt.Sprintf("where %s", strings.Join(whereFilter, " AND "))
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}
	return ph
}
