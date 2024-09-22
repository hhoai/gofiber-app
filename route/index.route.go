package route

import (
	"fiber-app/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	account := app.Group("/account", IsAuthenticated, CheckSession)

	app.Get("/information", IsAuthenticated, CheckSession, controller.InformationController)

	app.Post("/signup", controller.SignupPostController)

	app.Get("/logout", controller.LogoutController)

	app.Get("/login", controller.LoginController)

	app.Post("/login", controller.LoginPostController)

	app.Get("/admin", IsAuthenticated, CheckSession, controller.AdminController)

	account.Delete("/:id", controller.DeleteController).Name("deleteAccount")
	app.Delete("/delete-accounts", controller.DeleteMultipleAccounts)

	account.Get("/", CheckPermission, controller.CreateAccountController).Name("createAccount")

	account.Post("/", controller.CreateAccountPostController)

	account.Get("/:id", controller.UpdateAccountController).Name("updateAccount")

	account.Put("/:id", controller.UpdateAccountPutController)

	app.Get("/unauthorized", controller.Unauthorized)

	app.Get("/createRole", IsAuthenticated, CheckSession, CheckPermission, controller.CreateRoleController).Name("createRole")

	app.Post("/createRole", controller.CreateRolePostController)

	app.Get("/deleteRole/:id", controller.DeleteRoleController).Name("deleteRole")

	app.Get("/updateRole/:id", IsAuthenticated, CheckSession, controller.UpdateRoleController).Name("UpdateRole")

	app.Put("/updateRole/:id", controller.UpdateRolePutController)

	app.Get("/role", controller.RoleController)

	app.Get("/sales", controller.SalesController)

	app.Get("/data", controller.GetSalesData)
}
