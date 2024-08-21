package controller

import (
	"fiber-app/shared"

	"github.com/gofiber/fiber/v2"
)

func LogoutController(c *fiber.Ctx) error {
	sess, err := shared.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	sess.Destroy()
	return c.Redirect("/login")
}
