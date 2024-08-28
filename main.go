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

func main() {
	// initial database
	database.DatabaseInit()
	migration.RunMigration()

	controller.Store = session.New()

	engine := html.New("./views", ".html")
	engine.AddFuncMap(fiber.Map{
		"isPermissionSelected": controller.IsPermissionSelected,
	})

	app := fiber.New(fiber.Config{
		Views: engine, // Sử dụng template engine đã nạp
	})

	app.Static("/", "./public")

	// initial route
	route.RouteInit(app)

	app.Listen(":8080")
}

// import (
// 	"html/template"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/template/html/v2"
// 	"strings"
// )

// // Định nghĩa funcmap
// var funcMap = template.FuncMap{
// 	"uppercase": func(s string) string {
// 			return strings.ToUpper(s)
// 	},
// }

// func main() {
// 	app := fiber.New()

// 	// Cấu hình template engine với funcmap
// 	tmpl := html.New("./views", ".html").Funcs(funcMap)
// 	app.Renderer = tmpl

// 	// Định nghĩa một route ví dụ
// 	app.Get("/", func(c *fiber.Ctx) error {
// 			data := map[string]interface{}{
// 					"Name": "john",
// 			}
// 			return c.Render("index", data)
// 	})

// 	app.Listen(":3000")
// }
