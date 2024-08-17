package main

import (
	// "fmt"

	"fiber-app/database"
	"fiber-app/database/migration"
	"fiber-app/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

var store *session.Store

func main() {
	// initial database
	database.DatabaseInit()
	migration.RunMigration()

	engine := html.New("./views", ".html")

	store = session.New()

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	// initial route
	route.RouteInit(app)

	app.Listen(":8080")
}
