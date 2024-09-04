package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleController(c *fiber.Ctx) error {
	var permission []entity.Permission
	if err := database.DB.Find(&permission).Error; err != nil {
		log.Println(err)
	}

	sidebar := Sidebar(c)

	data := fiber.Map{
		"Title":        "Thêm vai trò",
		"Permission":   permission,
		"SidebarItems": sidebar,
	}

	return c.Render("createRole", data, "layouts/main")
}
