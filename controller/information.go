package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InformationController(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Kiểm tra người dùng đã đăng nhập chưa
		username := sess.Get("username")
		if username == nil {
			return c.Redirect("/login")
		}

		email := sess.Get("email")
		password := sess.Get("password")
		phone := sess.Get("phone")
		address := sess.Get("address")

		// Tạo dữ liệu để truyền vào template
		data := fiber.Map{
			"Username": username,
			"Email":    email,
			"Password": password,
			"Phone":    phone,
			"Address":  address,
		}
		return c.Render("information", data)
	}
}
