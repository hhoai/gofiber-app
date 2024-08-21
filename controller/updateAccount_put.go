package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateAccountPutController(c *fiber.Ctx) error {
	var p entity.Account
	id := c.Params("id")

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	nameError := validateUsername(p.Username)

	errorsMessage := fiber.Map{
		"NameError":     nameError,
		"PasswordError": validatePassword(p.Password),
		"EmailError":    validateEmail(p.Email),
		"AddressError":  validateAddress(p.Address),
		"PhoneError":    validatePhoneNumber(p.Phone),
	}

	var existingAccount entity.UserEntity
	// Kiểm tra sự tồn tại của tài khoản
	resultFindByUserName := database.DB.Where("name = ? ", p.Username).First(&existingAccount)
	resultFindByEmail := database.DB.Where("email = ?", p.Email).First(&existingAccount)

	if resultFindByUserName.Error != nil && resultFindByUserName.Error != gorm.ErrRecordNotFound {
		// Xử lý lỗi nếu có lỗi ngoài lỗi không tìm thấy
		return c.Status(fiber.StatusInternalServerError).JSON(resultFindByUserName.Error.Error())
	}

	if resultFindByEmail.Error != nil && resultFindByEmail.Error != gorm.ErrRecordNotFound {
		// Xử lý lỗi nếu có lỗi ngoài lỗi không tìm thấy
		return c.Status(fiber.StatusInternalServerError).JSON(resultFindByEmail.Error.Error())
	}

	if validateUsername(p.Username) != "" || validatePassword(p.Password) != "" || validateEmail(p.Email) != "" || validatePhoneNumber(p.Phone) != "" || validateAddress(p.Address) != "" {
		return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	}

	if resultFindByUserName.Error == nil {
		errorsMessage := fiber.Map{
			"NameError": "UserName already exists",
		}
		// Nếu không có lỗi và tìm thấy bản ghi, trả về lỗi trùng lặp
		return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	}

	if resultFindByEmail.Error == nil {
		errorsMessage := fiber.Map{
			"EmailError": "Email already exists",
		}
		// Nếu không có lỗi và tìm thấy bản ghi, trả về lỗi trùng lặp
		return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	}

	var account entity.UserEntity
	rs := database.DB.First(&account, "id = ?", id)

	if rs.Error != nil {
		log.Println(rs.Error)
	}

	account.Name = p.Username
	account.Email = p.Email
	account.Address = p.Address
	account.Phone = p.Phone

	if err := database.DB.Save(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var users []entity.UserEntity

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	user := fiber.Map{
		"Users": users,
	}
	return c.Render("admin", user, "layouts/main")
	// ...
}