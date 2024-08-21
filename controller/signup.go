package controller

import (
	"github.com/gofiber/fiber/v2"
)

func SignupController(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{
		"Title": "Đăng Kí",
	}, "layouts/main")
}
