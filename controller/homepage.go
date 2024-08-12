package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func HomepageController(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Trang Chủ",
		})
	}
}
