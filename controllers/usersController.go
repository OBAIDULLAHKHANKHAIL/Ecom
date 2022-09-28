package controllers

import (
	"fmt"
	"strconv"

	db "obaid/db"
	"obaid/models"
	model "obaid/models"

	"github.com/gofiber/fiber/v2"
)

func UserDetails(c *fiber.Ctx) error {

	param := c.Params("Id")

	var user model.User

	db.DB.Where("id=?", param).First(&user)
	if len(user.Name) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "Not Found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success!",
		"data":    user,
	})
}

func CreateUser(c *fiber.Ctx) error {

	fmt.Println("Creating a New User...")
	data := struct {
		Name string `json:"name"`
		User string `json:"user"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(504).JSON(fiber.Map{
			"success": false,
			"Message": "UnprocessiableEntity",
		})
	}
	if len(data.Name) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Bad Request",
			"error":   map[string]interface{}{},
		})
	}

	user := model.User{
		Name:      data.Name,
		User: data.User,
		Password: data.Password,
	}

	db.DB.Create(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User
	db.DB.Find(&user, "id = ?", userId)

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "User Not Found",
		})
	}

	var updateUserData models.User
	c.BodyParser(&updateUserData)
	if updateUserData.Name == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Empty Body"})
	}
	user.Name = updateUserData.Name
	db.DB.Table("Ecom").Where("id= ?", userId).Update("name", user.Name)
	db.DB.Table("Ecom").Where("id = ?",userId).Update("user", user.User)
	db.DB.Table("Ecom").Where("id = ?",userId).Update("password", user.Password)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	limit := c.Query("limit")
	skip := c.Query("skip")

	intLimit, _ := strconv.Atoi(limit)
	intSkip, _ := strconv.Atoi(skip)

	var user []model.User
	res := db.DB.Select([]string{"id", "name"}).Limit(intLimit).Offset(intSkip).Find(&user)

	fmt.Println(res.RowsAffected)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Sucess",
		"data":    user,
		"meta": map[string]interface{}{
			"total": res.RowsAffected,
			"limit": intLimit,
			"skip":  skip,
		},
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("Id")
	fmt.Println(id)
	user := model.User{}
	db.DB.First(&user, id)

	if user.Name == "" {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"Message": "No user found against this user id",
		})
	}

	result := db.DB.Delete(&user, "Id=?", id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "User Not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success!",
	})
}
