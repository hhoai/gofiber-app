package route

import (
	"fiber-app/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {

	// r.Get("/user", handler.UserHandlerGetAll)
	app.Get("/information", IsAuthenticated, CheckSession, controller.InformationController)

	app.Get("/signup", controller.SignupController)

	app.Post("/signup", controller.SignupPostController)

	app.Get("/logout", controller.LogoutController)

	app.Get("/login", controller.LoginController)

	app.Post("/login", controller.LoginPostController)

	app.Get("/admin", IsAuthenticated, CheckSession, controller.AdminController)

	app.Get("/deleteAccount/:id", IsAuthenticated, CheckSession, CheckPermission, controller.DeleteController).Name("deleteAccount")

	app.Get("/createAccount", IsAuthenticated, CheckSession, CheckPermission, controller.CreateAccountController).Name("createAccount")

	app.Post("/createAccount", controller.CreateAccountPostController)

	app.Get("/updateAccount/:id", IsAuthenticated, CheckSession, CheckPermission, controller.UpdateAccountController).Name("updateAccount")

	app.Put("/updateAccount/:id", controller.UpdateAccountPutController)

	app.Get("/unauthorized", controller.Unauthorized)

	app.Get("/permission", controller.PermissionController)

	app.Get("/createRole", controller.CreateRoleController)

	app.Post("/createRole", controller.CreateRolePostController)

	app.Get("/deleteRole/:id", controller.DeleteRoleController)

	app.Get("/updateRole/:id", IsAuthenticated, CheckSession, controller.UpdateRoleController)

	app.Put("/updateRole/:id", controller.UpdateRolePutController)

	app.Get("/role", controller.RoleController)
}
