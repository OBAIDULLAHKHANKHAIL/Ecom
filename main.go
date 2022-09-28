package main

import (
	"log"
	"os"

	// "fmt"
	// "obaid/db"
	"obaid/routes"

	"github.com/gofiber/fiber/v2"
)
// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, world")
}

func setupRoutes(app *fiber.App) {
	//checking
	app.Post("/hello", hello)

	//User Endpoints
	app.Post("/api/users", routes.CreateUser)
	// app.Get("/api/users", routes.GetUser)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Put("/api/users/:id", routes.UpdateUser)
	//Product Endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProduct)
	app.Get("/api/getAllProducts", routes.ListProduct)

	app.Delete("/api/DeleteProducts", routes.DeleteAll)

	app.Get("/api/products/:id", routes.GetProductById)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	app.Put("/api/products", routes.ListProduct)
}



func main() {

	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	app := fiber.New()
	
	setupRoutes(app)

	
	log.Fatal(app.Listen(":3000"))

	// db.Connect("","","","")
	// localhost:27017
	// Fiber instance
	// app = fiber.New()

	// 	initDB()

	// app.Get()

	// start server
	// app.Listen("3000")

	// defer db.DBConn.Close()
}
