package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"fiber-app/shared"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateAccountController(c *fiber.Ctx) error {
	sess, err := shared.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	role := sess.Get("role")

	if role == 2 {

		var p entity.UserEntity

		id := c.Params("id")

		result := database.DB.First(&p, "id = ?", id)

		if result.Error != nil {
			log.Println(result.Error)
		}

		// Tạo dữ liệu để truyền vào template
		data := fiber.Map{
			"id":       id,
			"Username": p.Name,
			"Email":    p.Email,
			"Phone":    p.Phone,
			"Address":  p.Address,
		}
		return c.Render("updateAccount", data, "layouts/main")

	}
	return c.Redirect("/login")
}
