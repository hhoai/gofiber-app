package controller

import (
	"fiber-app/shared"

	"github.com/gofiber/fiber/v2"
)

func CreateAccountController(c *fiber.Ctx) error {

	sess, err := shared.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	role := sess.Get("role")
	if role == 2 {
		return c.Render("createAccount", fiber.Map{
			"Title": "Tạo tài khoản",
		}, "layouts/main")
	}
	return c.Redirect("/login")

}
