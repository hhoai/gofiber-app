package controller

import "github.com/gofiber/fiber/v2"

func CreateAccountController(c *fiber.Ctx) error {
	return c.Render("createAccount", fiber.Map{
		"Title": "Tạo tài khoản",
	}, "layouts/main")
}
