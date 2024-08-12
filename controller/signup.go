package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SignupController(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{
			"Title": "Đăng Kí",
		})
	}
}
