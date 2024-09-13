package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PermissionUser(c *fiber.Ctx) []string {
	user := GetSessionUser(c)

	var RolePermission []entity.RolePermission
	if err := database.DB.Where("role_id = ?", user.RoleID).Find(&RolePermission); err != nil {
		log.Println(err)
	}

	var permissions []string
	for _, item := range RolePermission {
		var permission entity.Permission
		if err := database.DB.Where("id = ?", item.PermissionID).First(&permission).Error; err != nil {
			log.Println(err)
		} else {
			permissions = append(permissions, permission.Permission)
		}
	}

	log.Println("sidebar: ", permissions)

	return permissions
}
