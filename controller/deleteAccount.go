package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DeleteController(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	role := sess.Get("role")
	if role == 2 {
		var p entity.UserEntity

		id := c.Params("id")

		err := database.DB.Delete(&p, "id = ?", id).Error

		if err != nil {
			log.Println(err)
		}

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		username := sess.Get("username")

		log.Println(username, "delete")

		return c.Redirect("/admin")
	}

	return c.Redirect("/login")
}
