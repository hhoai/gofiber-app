package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"

	"log"

	"github.com/gofiber/fiber/v2"
)

func IsPermissionSelected(selectedPermissions []int, permissionID int) bool {
	for _, p := range selectedPermissions {
		if p == permissionID {
			return true
		}
	}
	return false
}

func UpdateRoleController(c *fiber.Ctx) error {
	var p entity.Role

	id := c.Params("id")

	result := database.DB.First(&p, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error)
	}

	var roles []entity.Role

	var permission []entity.Permission
	if err := database.DB.Find(&permission).Error; err != nil {
		log.Println(err)
	}

	rs := database.DB.Find(&roles)
	if rs.Error != nil {
		log.Println(rs.Error)
	}

	var rolPer []entity.RolePermission
	if err := database.DB.Where("role_id = ?", p.ID).Find(&rolPer).Error; err != nil {
		log.Println(err)
	}

	var accPer []entity.Permission
	var permissionID []int
	for _, p := range rolPer {
		if err := database.DB.Where("id = ?", p.PermissionID).Find(&accPer).Error; err != nil {
			log.Println(err)
		} else {
			permissionID = append(permissionID, p.PermissionID)
		}
	}

	// Tạo dữ liệu để truyền vào template
	data := fiber.Map{
		"id":           id,
		"Role":         p.Role,
		"Permission":   permission,
		"PermissionID": permissionID,
	}

	// tmpl := template.Must(template.New("").Funcs(template.FuncMap{
	// 	"isPermissionSelected": IsPermissionSelected,
	// }).ParseGlob("views/*.html"))

	// return tmpl.ExecuteTemplate(c.Response().BodyWriter(), "updateRole.html", data)

	return c.Render("updateRole", data, "layouts/main")
}
