package controller

import (
	"github.com/gofiber/fiber/v2"
)

func LogoutController(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	sess.Destroy()
	return c.Redirect("/login")
}
