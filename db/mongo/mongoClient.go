package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"obaid/config"
	"obaid/db"
	"obaid/models"
)

type DbInstance struct {
	DataB *mongo.Database
}

const (
	usersCollection    = "Users"
	productsCollection = "Products"
)

func init() {
	db.Register("mongo", NewMongoClient)
}

type client struct {
	conn *mongo.Client
}

// NewMongoClient initializes a mysql database connection.
func NewMongoClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	return &client{conn: cli}, nil
}

func (m client) AddUser(user *models.User) (string, error) {
	if user.ID == "" {
		return "id is empty", nil
	}
	user.ID = guuid.NewV4().String()

	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(usersCollection)
	if _, err := collection.InsertOne(context.TODO(), user); err != nil {
		return "", errors.Wrap(err, "failed to add user")
	}
	return user.ID, nil
}

func (m client) UpdateUser(id string, user *models.User) error {
	user.ID = id
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(usersCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": user.ID}}, bson.M{"$set": user}); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (m client) RemoveUserByID(id string) error {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(usersCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (m client) GetUserByID(id string) (*models.User, error) {
	var stu *models.User
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(usersCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&stu); err != nil {
		if mongo.ErrNoDocuments != nil {
			return nil, errors.Wrap(err, "failed to fetch user....not found")
		}
		return nil, err
	}
	return stu, nil
}

func (m client) AddProduct(product *models.Product) (string, error) {
	if product.ID == "" {
		return "id is not empty", nil
	}
	product.ID = guuid.NewV4().String()
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(productsCollection)
	if _, err := collection.InsertOne(context.TODO(), product); err != nil {
		return "", errors.Wrap(err, "failed to add product")
	}
	return product.ID, nil
}

func (m client) UpdateProduct(id string, product *models.Product) error {
	product.ID = id
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(productsCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": product.ID}}, bson.M{"$set": product}); err != nil {
		return errors.Wrap(err, "failed to update product")
	}
	return nil
}

func (m client) GetProductByID(id string) (*models.Product, error) {
	var stu *models.Product
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(productsCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&stu); err != nil {
		if mongo.ErrNoDocuments != nil {
			return nil, errors.Wrap(err, "failed to fetch product....not found")
		}
		return nil, err
	}
	return stu, nil
}

func (m client) RemoveProductByID(id string) error {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(productsCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete product")
	}

	return nil
}

func (m client) ListProduct(filter map[string]interface{}, lim int64, off int64) ([]*models.Product, error) {
	var std []*models.Product
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(productsCollection)
	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Skip:  &off,
		Limit: &lim,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve product")
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var em *models.Product
		if err = cursor.Decode(&em); err != nil {
			return nil, errors.Wrap(err, "failed to scan retrieved rows")
		}
		std = append(std, em)
	}
	return std, nil
}
