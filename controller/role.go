package controller

import (
	"fiber-app/database"

	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RoleController(c *fiber.Ctx) error {

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var user = GetSessionUser(c)

	log.Println(sess.Get("username"), "admin")

	if user.RoleID != 1 {
		var role []entity.RolePermissionWithRowNumber

		result := database.DB.Table("roles").
			Select("ROW_NUMBER() OVER (ORDER BY roles.id) AS RowNumber, roles.role AS role_name, id").
			Find(&role)

		if result.Error != nil {
			log.Println(result.Error)
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}

		log.Println(role)
		sidebar := Sidebar(c)

		return c.Render("role", fiber.Map{
			"Roles":        role,
			"SidebarItems": sidebar,
		}, "layouts/main")
	}
	return c.Redirect("/information")
}
