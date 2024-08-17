package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DeleteController(c *fiber.Ctx) error {
	var p entity.UserEntity

	id := c.Params("id")

	err := database.DB.Delete(&p, "id = ?", id).Error

	if err != nil {
		log.Println(err)
	}

	return c.Redirect("/admin")
}
