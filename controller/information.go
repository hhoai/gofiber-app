package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func InformationController(c *fiber.Ctx) error {

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
	phone := sess.Get("phone")
	address := sess.Get("address")

	log.Println(address)
	// Tạo dữ liệu để truyền vào template
	data := fiber.Map{
		"Username": username,
		"Email":    email,
		"Phone":    phone,
		"Address":  address,
	}
	return c.Render("information", data, "layouts/main")
}
