package routes

import (
	"errors"
	"fmt"
	"obaid/db"
	"obaid/db/mongo"
	"obaid/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	// this is not the model Product, see this as the serializer
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Password    string    `json:"password"`
	CreatedBy   string    `json:"createdby"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedBy   string    `json:"updatedby"`
	UpdatedAt   time.Time `json:"updatedat"`
}



func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	client, err := mongo.NewMongoClient(db.Option{})
	// mysql.NewMysqlClient(db.Option{})
	if err != nil {
		panic(err)
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// db.Database.Db.Create(&product)
	_, err = client.AddProduct(&product)
	if err != nil {
		panic(err)
	}
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return c.Status(200).JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	product := []models.Product{}

	// db.Database.Db.Find(&product)
	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}
	responseProducts := []Product{}
	for _, product := range product {

		responseProducts = append(responseProducts, Product(product))
		_, err = client.GetProductByID(product.ID)
		if err != nil {
			panic(err)
		}
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id string, product *models.Product) error {
	// db.Database.Db.Find(&product, "id=?", id)
	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}

	client.GetProductByID(id)
	if product.ID == "0" {
		return errors.New("Product does not exist")
	}
	return nil
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var product *models.Product

	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}

	product, err = client.GetProductByID(id)
	if err != nil {
		panic(err)
	}

	return c.Status(200).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}

	type UpdateProduct struct {
		Name        string    `json:"name"`
		Price       int       `json:"price"`
		Description string    `json:"description"`
		Password    string    `json:"password"`
		CreatedBy   string    `json:"createdby"`
		CreatedAt   time.Time `json:"createdat"`
		UpdatedBy   string    `json:"updatedby"`
		UpdatedAt   time.Time `json:"updatedat"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	fmt.Println(updateData.Name)
	fmt.Println(product.Name)
	product.Name = updateData.Name
	product.Price = updateData.Price
	product.Description = updateData.Description
	product.Password = updateData.Password
	product.CreatedBy = updateData.CreatedBy
	product.CreatedAt = time.Now()
	product.UpdatedBy = updateData.UpdatedBy
	product.UpdatedAt = updateData.UpdatedAt
	product.UpdatedAt = time.Now()

	fmt.Println("hi")

	client.UpdateProduct(id, &product)
	if err != nil {
		panic(err)
	}

	return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}

	if err := client.RemoveProductByID(id); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Sucessfully Deleted Product")

}

func DeleteAll(c *fiber.Ctx) error {
	// product := []models.Product{}

	// client, err := mongo.NewMongoClient(db.Option{})
	// if err != nil {
	// 	panic(err)
	// }

	// db.Database.Db.Delete(product)

	return nil
}

func ListProduct(c *fiber.Ctx) error {
	products := []models.Product{}

	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		client.GetUserByID(product.ID)
		// _ = db.Database.Db.Find(&product)
		fmt.Println(product)
	}

	return c.Status(200).JSON(products)
}