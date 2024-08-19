package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginPostController(c *fiber.Ctx) error {
	var p entity.Account

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	var existingAccount entity.UserEntity
	// Kiểm tra sự tồn tại của tài khoản
	result := database.DB.First(&existingAccount, "name = ?", p.Username)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		// Xử lý lỗi nếu có lỗi ngoài lỗi không tìm thấy
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	if result.Error != nil {
		if entity.CheckPassword(&existingAccount, p.Password) {
			// Không tìm thấy bản ghi, trả về lỗi đăng nhập thất bại
			return c.Status(fiber.StatusUnauthorized).SendString("Incorrect username or password")
		} else {
			// Lỗi khác, trả về lỗi server
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}
	}

	result = database.DB.First(&existingAccount, existingAccount.Name)

	// Lưu thông tin người dùng vào session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sess.Set("username", existingAccount.Name)
	sess.Set("email", existingAccount.Email)
	sess.Set("password", existingAccount.Password)
	sess.Set("phone", existingAccount.Phone)
	sess.Set("address", existingAccount.Address)
	sess.Set("role", existingAccount.Role)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if existingAccount.Role == 2 {
		return c.Redirect("/admin")
	}
	return c.Redirect("/information")
}
