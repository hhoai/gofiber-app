package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func HomepageController(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": "Trang Chủ",
	}, "layouts/main")
}
