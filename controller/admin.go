package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AdminController(c *fiber.Ctx) error {
	var users []entity.UserEntity
	// Truy vấn tất cả các bản ghi
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	user := fiber.Map{
		"Users": users,
	}

	return c.Render("admin", user, "layouts/main")
}
