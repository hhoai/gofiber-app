package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func LoginController(c *fiber.Ctx) error {
	// Kiểm tra người dùng đã đăng nhập chưa
	store = session.New()

	sess, err := store.Get(c)
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
