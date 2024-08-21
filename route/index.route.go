package route

import (
	"fiber-app/controller"
	"fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	locations := app.Group("/locations", middleware.IsAuthenticated, middleware.CheckSession)
	app.Get("/", controller.HomepageController)
	// r.Get("/user", handler.UserHandlerGetAll)
	app.Get("/information", controller.InformationController)

	app.Get("/logout", controller.LogoutController)

	app.Get("/login", controller.LoginController)
	app.Get("/signup", controller.SignupController)
	app.Post("/signup", controller.SignupPostController)

	app.Post("/login", controller.LoginPostController)

	app.Get("/admin", controller.AdminController)

	app.Get("/deleteAccount/:id", controller.DeleteController)

	app.Get("/createAccount", controller.CreateAccountController)

	app.Post("/createAccount", controller.CreateAccountPostController)

	locations.Get("/updateAccount/:id", controller.UpdateAccountController)

	locations.Put("/updateAccount/:id", controller.UpdateAccountPutController)
}
