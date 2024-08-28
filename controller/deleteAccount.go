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

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	username := sess.Get("username")

	log.Println(username, "delete")

	return c.Redirect("/admin")
}
