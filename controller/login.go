package controller

import (
	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	// Kiểm tra người dùng đã đăng nhập chưa

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	username := sess.Get("username")
	if username != nil {
		return c.Redirect("/information")
	}

	return c.Render("login", fiber.Map{
		"Title": "Đăng Nhập",
	}, "layouts/main")
}
