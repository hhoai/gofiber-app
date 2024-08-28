package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateAccountPutController(c *fiber.Ctx) error {
	id := c.Params("id")
	var p entity.Account

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	nameError := validateUsername(p.Username)

	errorsMessage := fiber.Map{
		"NameError":    nameError,
		"EmailError":   validateEmail(p.Email),
		"AddressError": validateAddress(p.Address),
		"PhoneError":   validatePhoneNumber(p.Phone),
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

	if validateUsername(p.Username) != "" || validateEmail(p.Email) != "" || validatePhoneNumber(p.Phone) != "" || validateAddress(p.Address) != "" {
		return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	}

	// if resultFindByUserName.Error == nil {
	// 	errorsMessage := fiber.Map{
	// 		"NameError": "UserName already exists",
	// 	}
	// 	// Nếu không có lỗi và tìm thấy bản ghi, trả về lỗi trùng lặp
	// 	return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	// }

	// if resultFindByEmail.Error == nil {
	// 	errorsMessage := fiber.Map{
	// 		"EmailError": "Email already exists",
	// 	}
	// 	// Nếu không có lỗi và tìm thấy bản ghi, trả về lỗi trùng lặp
	// 	return c.Status(fiber.StatusConflict).JSON(errorsMessage)
	// }

	var account entity.UserEntity
	rs := database.DB.First(&account, "id = ?", id)

	if rs.Error != nil {
		log.Println(rs.Error)
	}

	account.Name = p.Username
	account.Email = p.Email
	account.Address = p.Address
	account.Phone = p.Phone
	account.RoleID = p.RoleID

	if err := database.DB.Save(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	username := sess.Get("username")

	log.Println(username, "update")
	var users []entity.UserEntity

	err1 := database.DB.Find(&users).Error
	if err1 != nil {
		log.Println(err1)
	}

	user := fiber.Map{
		"Users": users,
	}
	return c.Render("admin", user, "layouts/main")
	// return c.Redirect("/admin")
	// ...
}
