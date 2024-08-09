package main

import (
	// "fmt"
	"fiber-app/database"
	"fiber-app/database/migration"
	"fiber-app/model/entity"
	"fiber-app/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"

	// "fmt"
	"regexp"
	"strings"
)

var store *session.Store

func validateUsername(username string) string {
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

func main() {
	// initial database
	database.DatabaseInit()
	migration.RunMigration()

	engine := html.New("./view", ".html")

	store = session.New()

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home Page",
		})
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		sess.Destroy()
		return c.Redirect("/")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		// Kiểm tra người dùng đã đăng nhập chưa
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		username := sess.Get("username")
		if username != nil {
			return c.Redirect("/information")
		}

		return c.Render("login", fiber.Map{
			"Title": "Login Page",
		})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{
			"Title": "Signup Page",
		})
	})

	app.Get("/information", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Kiểm tra người dùng đã đăng nhập chưa
		username := sess.Get("username")
		if username == nil {
			return c.Redirect("/login")
		}

		email := sess.Get("email")
		password := sess.Get("password")
		phone := sess.Get("phone")
		address := sess.Get("address")

		// Tạo dữ liệu để truyền vào template
		data := fiber.Map{
			"Username": username,
			"Email":    email,
			"Password": password,
			"Phone":    phone,
			"Address":  address,
		}
		return c.Render("information", data)
	})

	type Account struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Address  string `json:"address"`
		Phone    string `json:"phone"`
	}

	app.Post("/signup", func(c *fiber.Ctx) error {
		var p Account

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

		// Tạo bảng nếu chưa tồn tại
		database.DB.AutoMigrate(&Account{})

		var existingAccount entity.UserEntity
		// Kiểm tra sự tồn tại của tài khoản
		result := database.DB.Where("name = ? OR email = ?", p.Username, p.Email).First(&existingAccount)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			// Xử lý lỗi nếu có lỗi ngoài lỗi không tìm thấy
			return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
		}

		if result.Error == nil {
			// Nếu không có lỗi và tìm thấy bản ghi, trả về lỗi trùng lặp
			return c.Status(fiber.StatusConflict).JSON("Username or Email already exists")
		}

		if validateUsername(p.Username) != "" || validatePassword(p.Password) != "" || validateEmail(p.Email) != "" || validatePhoneNumber(p.Phone) != "" || validateAddress(p.Address) != "" {
			return c.Status(fiber.StatusConflict).JSON(errorsMessage)
		}
		// Tạo tài khoản mới
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
		sess, err := store.Get(c)
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
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var p Account

		if err := c.BodyParser(&p); err != nil {
			return err
		}

		// Tạo bảng nếu chưa tồn tại
		database.DB.AutoMigrate(&Account{})

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

		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Redirect("/information")
	})

	// initial route
	route.RouteInit(app)

	app.Listen(":8080")
}
