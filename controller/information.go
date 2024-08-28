package controller

import (
	"github.com/gofiber/fiber/v2"
)

func InformationController(c *fiber.Ctx) error {

	user := GetSessionUser(c)

	// log.Println(address)
	// Tạo dữ liệu để truyền vào template
	data := fiber.Map{
		"Username": user.Name,
		"Email":    user.Email,
		"Phone":    user.Phone,
		"Address":  user.Address,
	}
	return c.Render("information", data, "layouts/main")
}
