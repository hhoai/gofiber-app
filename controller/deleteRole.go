package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DeleteRoleController(c *fiber.Ctx) error {
	var p entity.Role
	var rolPer entity.RolePermission

	id := c.Params("id")

	if err := database.DB.Delete(&p, "id = ?", id).Error; err != nil {
		log.Println(err)
	}

	if err := database.DB.Delete(&rolPer, "role_id = ?", id).Error; err != nil {
		log.Println(err)
	}

	return c.Redirect("/role")
}
