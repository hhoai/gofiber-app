package main

import (
	// "fmt"
	"fiber-app/database"
	"fiber-app/database/migration"
	"fiber-app/route"
	"fiber-app/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// initial database
	database.DatabaseInit()
	migration.RunMigration()

	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	shared.InitSession()
	// initial route
	route.RouteInit(app)

	app.Listen(":8080")
}
