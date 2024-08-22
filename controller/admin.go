// package controller

// import (
// 	"fiber-app/database"
// 	"fiber-app/model/entity"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// )

// func AdminController(c *fiber.Ctx) error {

// 	sess, err := store.Get(c)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	role := sess.Get("role")

// 	if role == 2 {
// 		var users []entity.UserEntity
// 		// Truy vấn tất cả các bản ghi
// 		result := database.DB.Find(&users)
// 		if result.Error != nil {
// 			log.Println(result.Error)
// 		}

// 		err := database.DB.Find(&users).Error
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		user := fiber.Map{
// 			"Users": users,
// 		}
// 		return c.Render("admin", user, "layouts/main")
// 	}
// 	return c.Redirect("/information")
// }

package controller

import (
	"fiber-app/database"

	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

// func GetSessionUser (c fiber.Ctx) entity.UserEntity {
// 	var user entity.UserEntity
// 	var rolPer[] entity.Permission
// 	sess,_ := session.Get(c)
// 	username := sess. Get ("username")
// 	if err := database.DB.Where("BINARY username = ?", username).First(&user). Error; err != nil {
// 	database.OutputDebug("username not found in session")
// 	}
// 	if err:= database.DB.Model(&entity.RolePermission{}).
// 	Joins("Permission").
// 	Where("role_permission.role_id", user.RoleID).
// 	Where("role_permission.permission_id")
// 	database.OutputDebug("permission not found in check session"){
// 	}
// 	sess.Set("rolePermission", rolPer)
// 	if err := sess.Save(); err != nil {
// 	}
// 	initializers. OutputDebug("Can not save session role permission in check session")
// 	return user
// 	}

func AdminController(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	role := sess.Get("role")

	log.Println(sess.Get("username"), "admin")

	if role == 2 {
		var users []entity.UserWithRowNumber

		// Truy vấn SQL để lấy RowNumber
		// sqlQuery := `
		// 	SELECT ROW_NUMBER() OVER (ORDER BY id) AS RowNumber, id, name, email, address, phone, role_id
		// 	FROM user_entities
		// `

		// sqlQuery := `SELECT ROW_NUMBER() OVER (ORDER BY u.id) AS RowNumber, u.id, u.name, u.email, u.address, u.phone, r.role AS RoleName
		// 	FROM user_entities u
		// 	JOIN roles r ON u.role_id = r.id `

		result := database.DB.Table("user_entities").
			Joins("INNER JOIN roles ON user_entities.role_id = roles.id").
			Select("ROW_NUMBER() OVER (ORDER BY id) AS RowNumber, user_entities.id, name, email, address, phone, roles.role AS role_name").
			Find(&users)

		// result := database.DB.Raw(sqlQuery).Scan(&users)
		if result.Error != nil {
			log.Println(result.Error)
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}

		user := fiber.Map{
			"Users": users,
		}

		log.Println(user)

		return c.Render("admin", user, "layouts/main")
	}
	return c.Redirect("/information")
}
