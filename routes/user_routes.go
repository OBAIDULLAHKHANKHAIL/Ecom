package routes

import (
	"errors"
	"fmt"
	"obaid/db"
	"obaid/db/mongo"
	"obaid/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	// this is not the model User, see this as the serializer
	ID       string `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// func CreateResponserUser(userModel models.User) User {
// 	return User{ID: userModel.ID, Name: userModel.Name, User: userModel.User, Password: userModel.Password}
// }

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	fmt.Println("Create User Called...")
	client, err := mongo.NewMongoClient(db.Option{})
	// mysql.NewMysqlClient(db.Option{})
	if err != nil {
		panic(err)
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// responseUser := CreateResponserUser(user)
	_, err = client.AddUser(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(&user)
	// db.Database.Db.Create(&user)
	return c.Status(200).JSON(user)
}

// func GetUser(c *fiber.Ctx) error {
// 	users := []models.User{}

// 	client, err := mongo.NewMongoClient(db.Option{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// db.Database.Db.Find(&users)

// 	client.

// 	responseUsers := []User{}
// 	for _, user := range users {
// 		responseUser := CreateResponserUser(user)
// 		responseUsers = append(responseUsers, responseUser)
// 	}
// 	return c.Status(200).JSON(responseUsers)
// }

func findUser(id string, user *models.User) error {
	// db.Database.Db.Find(&user, "id=?", id)
	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}
	client.GetUserByID(id)
	if user.ID == "0" {
		return errors.New("User does not exist")
	}
	return nil
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user *models.User
	// fmt.Println(id)
	// if err != nil {
	// 	return c.Status(400).JSON("Please ensure that :id is an integer")
	// }

	// mysql.NewMysqlClient(db.Option{})

	// if err := findUser(id, &user); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}
	user, err = client.GetUserByID(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(&user)

	// responseUser := CreateResponserUser(&user)
	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// if err != nil {
	// 	return c.Status(400).JSON("Please ensure that :id is an integer")
	// }

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.Name = updateData.Name
	user.User = updateData.User
	user.Password = updateData.Password

	// db.Database.Db.Save(&user)

	client, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		panic(err)
	}
	client.UpdateUser(id, &user)
	fmt.Println(user)
	// responseUser := CreateResponserUser(user)
	return c.Status(200).JSON(user)
}
