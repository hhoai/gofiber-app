package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateRolePutController(c *fiber.Ctx) error {
	id := c.Params("id")

	var data entity.PermissionsRequest
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var inputRole entity.Role
	if err := c.BodyParser(&inputRole); err != nil {
		return err
	}

	// Lấy giá trị từ form
	permissions := data.Permissions

	log.Println("permissions", permissions)

	var updateRole entity.Role

	if err := database.DB.Where("id = ?", id).First(&updateRole).Error; err != nil {
		log.Println(err)
	}

	updateRole.Role = inputRole.Role
	if err := database.DB.Save(&updateRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var rolePermission entity.RolePermission
	if err := database.DB.Where("role_id = ?", updateRole.ID).Delete(&rolePermission).Error; err != nil {
		log.Println(err)
	}

	for _, permission := range permissions {
		log.Println("permission: ", permission)
		rolPer := entity.RolePermission{
			RoleID:       updateRole.ID,
			PermissionID: permission,
		}
		if err := database.DB.Create(rolPer).Error; err != nil {
			log.Println(err)
		}
		log.Println(rolPer)
	}

	var role []entity.RolePermissionWithRowNumber

	result := database.DB.Table("roles").
		Select("ROW_NUMBER() OVER (ORDER BY roles.id) AS RowNumber, roles.role AS role_name, id").
		Find(&role)

	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	log.Println(role)
	// roles :=

	return c.Render("role", fiber.Map{
		"Roles": role,
	}, "layouts/main")
}
