package controller

import (
	"fiber-app/database"

	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IsPermissionSelected(selectedPermissions []int, permissionID int) bool {
	for _, p := range selectedPermissions {
		if p == permissionID {
			return true
		}
	}
	return false
}

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

		return c.Render("role", fiber.Map{
			"Roles": role,
			"Ctx":   c,
		}, "layouts/main")
	}
	return c.Redirect("/information")
}

// create role
func CreateRoleController(c *fiber.Ctx) error {
	var permission []entity.Permission
	if err := database.DB.Find(&permission).Error; err != nil {
		log.Println(err)
	}

	data := fiber.Map{
		"Title":      "Thêm vai trò",
		"Permission": permission,
		"Ctx":        c,
	}

	return c.Render("createRole", data, "layouts/main")
}

// handle create role
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
	if err := database.DB.Where("role = ?", p.Role).First(&newRole).Error; err != nil {
		log.Println(err)
	}

	for _, permission := range permissions {
		log.Println("permission: ", permission)
		rolPer := entity.RolePermission{
			RoleID:       newRole.ID,
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
	// ...
}

// update role
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
		"Ctx":          c,
	}

	// tmpl := template.Must(template.New("").Funcs(template.FuncMap{
	// 	"isPermissionSelected": IsPermissionSelected,
	// }).ParseGlob("views/*.html"))

	// return tmpl.ExecuteTemplate(c.Response().BodyWriter(), "updateRole.html", data)

	return c.Render("updateRole", data, "layouts/main")
}

// handle update role

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

// delete role

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
