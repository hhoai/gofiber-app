package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateAccountController(c *fiber.Ctx) error {
	var roles []entity.Role

	result := database.DB.Find(&roles)
	if result.Error != nil {
		log.Println(result.Error)
	}

	sidebar := Sidebar(c)

	data := fiber.Map{
		"Role":         roles,
		"SidebarItems": sidebar,
	}

	return c.Render("createAccount", data, "layouts/main")
}
