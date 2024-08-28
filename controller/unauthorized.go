package controller

import "github.com/gofiber/fiber/v2"

func Unauthorized(c *fiber.Ctx) error {
	return c.Render("unauthorized", fiber.Map{}, "layouts/main")
}
