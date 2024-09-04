package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateAccountController(c *fiber.Ctx) error {
	var p entity.UserEntity

	id := c.Params("id")

	result := database.DB.First(&p, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error)
	}

	var roles []entity.Role

	rs := database.DB.Find(&roles)
	if rs.Error != nil {
		log.Println(rs.Error)
	}

	sidebar := Sidebar(c)

	// Tạo dữ liệu để truyền vào template
	data := fiber.Map{
		"id":           id,
		"Username":     p.Name,
		"Email":        p.Email,
		"Phone":        p.Phone,
		"Address":      p.Address,
		"RoleID":       p.RoleID,
		"SidebarItems": sidebar,
		"Roles":        roles,
	}
	return c.Render("updateAccount", data, "layouts/main")
}
