package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"fiber-app/shared"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// "fmt"
	"regexp"
	"strings"
)

func validateUsername(username string) string {
	//store = session.New()

	if strings.TrimSpace(username) == "" {
		return "Username must not be empty!"
	} else {
		if len(strings.TrimSpace(username)) < 8 || len(strings.TrimSpace(username)) > 30 {
			return "Username must be between 8 and 30 characters!"
		}
	}
	return ""
}

func validateEmail(email string) string {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	isValid := re.MatchString(email)

	if strings.TrimSpace(email) == "" {
		return "Email must not be empty!"
	} else {
		if !isValid {
			return "Email is not a valid!"
		}
	}
	return ""
}

func validatePhoneNumber(phoneNumber string) string {
	regex := `^0[0-9]{9}$`
	re := regexp.MustCompile(regex)
	isValid := re.MatchString(phoneNumber)

	if strings.TrimSpace(phoneNumber) == "" {
		return "Phone number must not be empty!"
	} else {
		if !isValid {
			return "Phone number is not a valid!"
		}
	}
	return ""
}

func validatePassword(password string) string {
	if strings.TrimSpace(password) == "" {
		return "Password must not be empty!"
	} else {
		if len(strings.TrimSpace(password)) < 8 || len(strings.TrimSpace(password)) > 20 {
			return "Password must be between 8 and 20 characters!"
		}
	}
	return ""
}

func validateAddress(address string) string {
	if strings.TrimSpace(address) == "" {
		return "Address must not be empty!"
	}
	return ""
}
func SignupPostController(c *fiber.Ctx) error {
	var p entity.Account

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
	// Tạo tài khoản mới
	log.Println(p.Address)
	newAccount := entity.UserEntity{
		Name:     p.Username,
		Password: p.Password,
		Email:    p.Email,
		Address:  p.Address,
		Phone:    p.Phone,
	}

	if err := database.DB.Create(&newAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	// Lưu thông tin người dùng vào session
	sess, err := shared.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	sess.Set("username", p.Username)
	sess.Set("email", p.Email)
	sess.Set("password", p.Password)
	sess.Set("phone", p.Phone)
	sess.Set("address", p.Address)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Redirect("/information")
	// ...
}
