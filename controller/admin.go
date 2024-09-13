package controller

import (
	"fiber-app/database"
	"strings"

	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func GetSessionUser(c *fiber.Ctx) entity.UserEntity {
	var user entity.UserEntity
	var rolPer []entity.RolePermission

	sess, _ := Store.Get(c)
	username := sess.Get("username")

	if err := database.DB.Where("name = ?", username).First(&user).Error; err != nil { // binary: phan biet chu hoa, chu thuong
		log.Println("username not found in session")
	}

	// if err := database.DB.Model(&entity.RolePermission{}).Joins("INNER JOIN Permission ON role_permissions.permission_id = permissions.id").
	// 	Select("permission_id").Where("role_id = ?", user.RoleID).
	// 	Find(&rolPer).Error; err != nil {
	// 	log.Println("permission not found in session")
	// }

	sess.Set("rolPermission", rolPer)

	return user
}

func CheckPermissionUser(s string, c *fiber.Ctx) bool {
	permissions := PermissionUser(c)
	for _, p := range permissions {
		if strings.EqualFold(p, s) {
			return true
		}
	}
	return false
}

func AdminController(c *fiber.Ctx) error {

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	log.Println(sess.Get("username"), "admin")

	// if user.RoleID != 1 {
	// }
	var users []entity.UserWithRowNumber

	result := database.DB.Table("user_entities").
		Joins("INNER JOIN roles ON user_entities.role_id = roles.id").
		Select("ROW_NUMBER() OVER (ORDER BY id) AS RowNumber, user_entities.id, name, email, address, phone, roles.role AS role_name").
		Find(&users)

	// result := database.DB.Raw(sqlQuery).Scan(&users)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	data := fiber.Map{
		"Users": users,
		"Ctx":   c,
	}

	// log.Println(user)

	return c.Render("admin", data, "layouts/main")
}
