package routes

import (
	CashierController "obaid/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routings(app *fiber.App) {

	app.Post("/user", CashierController.CreateUser)
	app.Get("/user/:Id", CashierController.UserDetails)
	app.Put("/user/:Id", CashierController.UpdateUser)
	app.Get("/user", CashierController.GetUsers)
	app.Delete("/user/:Id", CashierController.DeleteUser)

}