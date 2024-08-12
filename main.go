package main

import (
	// "fmt"

	"fiber-app/controller"
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

	app.Get("/", controller.HomepageController(store))

	app.Get("/logout", controller.LogoutController(store))

	app.Get("/login", controller.LoginController(store))

	app.Get("/signup", controller.SignupController(store))

	app.Get("/information", controller.InformationController(store))

	app.Post("/signup", controller.SignupPostController(store))

	app.Post("/login", controller.LoginPostController(store))

	// initial route
	route.RouteInit(app)

	app.Listen(":8080")
}
