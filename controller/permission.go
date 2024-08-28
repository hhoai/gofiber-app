package controller

import (
	"fiber-app/database"

	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PermissionController(c *fiber.Ctx) error {

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var user = GetSessionUser(c)

	log.Println(sess.Get("username"), "admin")

	if user.RoleID != 1 {
		var permissions []entity.PermissionWithRowNumber

		result := database.DB.Table("permissions").
			Select("ROW_NUMBER() OVER (ORDER BY id) AS RowNumber, id, permission").
			Find(&permissions)

		if result.Error != nil {
			log.Println(result.Error)
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}

		user := fiber.Map{
			"Permissions": permissions,
		}

		return c.Render("permission", user, "layouts/main")
	}
	return c.Redirect("/information")
}
