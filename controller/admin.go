package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"fiber-app/shared"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUserSessionNow(c *fiber.Ctx) entity.UserEntity {
	var user entity.UserEntity
	sess, _ := shared.Store.Get(c)
	var username = sess.Get("username")
	if err := database.DB.Where("name = ?", username).First(&user).Error; err != nil {
		log.Println("username not found in session")
	}
	return user
}

func AdminController(c *fiber.Ctx) error {

	sess, err := shared.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	role := sess.Get("role")

	if role == 2 {
		var users []entity.UserEntity
		// Truy vấn tất cả các bản ghi
		result := database.DB.Find(&users)
		if result.Error != nil {
			log.Println(result.Error)
		}

		err := database.DB.Find(&users).Error
		if err != nil {
			log.Println(err)
		}

		user := fiber.Map{
			"Users": users,
		}
		return c.Render("admin", user, "layouts/main")
	}
	return c.Redirect("/information")
}
