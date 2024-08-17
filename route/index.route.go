package route

import (
	"fiber-app/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {

	// r.Get("/user", handler.UserHandlerGetAll)
	app.Get("/information", controller.InformationController)

	app.Get("/logout", controller.LogoutController)

	app.Get("/login", controller.LoginController)

	app.Post("/login", controller.LoginPostController)

	app.Get("/admin", controller.AdminController)

	app.Get("/deleteAccount/:id", controller.DeleteController)

	app.Get("/createAccount", controller.CreateAccountController)

	app.Post("/createAccount", controller.CreateAccountPostController)

	app.Get("/updateAccount/:id", controller.UpdateAccountController)

	app.Put("/updateAccount/:id", controller.UpdateAccountPutController)
}
