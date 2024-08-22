package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateAccountController(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	role := sess.Get("role")
	log.Println(sess.Get("username"), "create account view")
	var roles []entity.Role

	result := database.DB.Find(&roles)
	if result.Error != nil {
		log.Println(result.Error)
	}

	data := fiber.Map{
		"Role": roles,
	}

	if role == 2 {
		return c.Render("createAccount", data, "layouts/main")
	}
	return c.Redirect("/login")

}
