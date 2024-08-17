package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func HomepageController(c *fiber.Ctx) error {
	store = session.New()

	return c.Render("index", fiber.Map{
		"Title": "Trang Chủ",
	}, "layouts/main")
}
