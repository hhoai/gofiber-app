package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// create account
func CreateAccountController(c *fiber.Ctx) error {
	var roles []entity.Role

	result := database.DB.Find(&roles)
	if result.Error != nil {
		log.Println(result.Error)
	}

	data := fiber.Map{
		"Role": roles,
		"Ctx":  c,
	}

	return c.Render("createAccount", data, "layouts/main")
}

// post
func CreateAccountPostController(c *fiber.Ctx) error {
	var p entity.Account

	if err := c.BodyParser(&p); err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println(p.RoleID)
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
		RoleID:   p.RoleID,
	}

	if err := database.DB.Create(&newAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	username := sess.Get("username")

	log.Println(username, "create", p.RoleID)

	return c.Redirect("/admin")
	// ...
}

// update account
func UpdateAccountController(c *fiber.Ctx) error {
	var p entity.UserEntity

	id := c.Params("id")

	result := database.DB.First(&p, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error)
	}

	var roles []entity.Role

	rs := database.DB.Find(&roles)
	if rs.Error != nil {
		log.Println(rs.Error)
	}

	// Tạo dữ liệu để truyền vào template
	data := fiber.Map{
		"id":       id,
		"Username": p.Name,
		"Email":    p.Email,
		"Phone":    p.Phone,
		"Address":  p.Address,
		"RoleID":   p.RoleID,
		"Roles":    roles,
		"Ctx":      c,
	}
	return c.Render("updateAccount", data, "layouts/main")
}

// put
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
		"Ctx":   c,
	}
	return c.Render("admin", user, "layouts/main")
	// return c.Redirect("/admin")
	// ...
}

// delete

func DeleteController(c *fiber.Ctx) error {
	var p entity.UserEntity

	id := c.Params("id")

	err := database.DB.Delete(&p, "id = ?", id).Error

	if err != nil {
		log.Println(err)
	}

	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	username := sess.Get("username")

	log.Println(username, "delete")

	var users []entity.UserEntity

	err1 := database.DB.Find(&users).Error
	if err1 != nil {
		log.Println(err1)
	}

	user := fiber.Map{
		"Users": users,
		"Ctx":   c,
	}
	return c.Render("admin", user, "layouts/main")
}

func DeleteMultipleAccounts(c *fiber.Ctx) error {
	// Định nghĩa một struct để nhận mảng ID từ client
	type RequestBody struct {
		AccountIDs []int `json:"account_id"`
	}

	var reqBody RequestBody

	// Parse JSON từ request body
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dữ liệu không hợp lệ",
		})
	}
	// Kiểm tra nếu không có account ID nào được gửi lên
	if len(reqBody.AccountIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Vui lòng chọn bản ghi cần xóa",
		})
	}

	// Xóa nhiều bản ghi trong bảng user_entities dựa vào mảng ID
	if err := database.DB.Where("id IN ?", reqBody.AccountIDs).Delete(&entity.UserEntity{}).Error; err != nil {
		log.Println("Đã xảy ra lỗi:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Đã xảy ra lỗi",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Xóa thành công",
	})
}
