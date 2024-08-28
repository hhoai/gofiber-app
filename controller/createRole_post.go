package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRolePostController(c *fiber.Ctx) error {
	var p entity.Role

	if err := c.BodyParser(&p); err != nil {
		log.Println(err.Error())
		return err
	}

	var existingRole entity.Role
	// Kiểm tra sự tồn tại của tài khoản
	resultFindByRoleName := database.DB.Where("role = ?", p.Role).First(&existingRole)

	if resultFindByRoleName.Error != nil && resultFindByRoleName.Error != gorm.ErrRecordNotFound {
		// Xử lý lỗi nếu có lỗi ngoài lỗi không tìm thấy
		return c.Status(fiber.StatusInternalServerError).JSON(resultFindByRoleName.Error.Error())
	}

	if resultFindByRoleName.Error == nil {
		errorsMessage := fiber.Map{
			"RoleNameError": "Role already exists",
		}
		// Nếu không có lỗi và tìm thấy bản ghi, trả về lỗi trùng lặp
		return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	}

	if err := database.DB.Create(&p).Error; err != nil {
		log.Println(err)
	}

	var data entity.PermissionsRequest
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Lấy giá trị từ form
	permissions := data.Permissions

	log.Println("permissions", permissions)

	var newRole entity.Role
	if err := database.DB.Where("role = ?", p.Role).Find(&newRole).Error; err != nil {
		log.Println(err)
	}

	for _, permission := range permissions {
		log.Println("permission: ", permission)
		var per entity.Permission
		if err := database.DB.Where("permission = ?", permission).Find(&per).Error; err != nil {
			log.Println(err)
		} else {
			rolPer := entity.RolePermission{
				RoleID:       newRole.ID,
				PermissionID: per.ID,
			}
			if err := database.DB.Create(rolPer).Error; err != nil {
				log.Println(err)
			}
			log.Println(rolPer)
		}
	}

	log.Println("create", p.Role)

	return c.Redirect("/role")
	// ...
}
