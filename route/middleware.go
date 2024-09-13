package route

import (
	"fiber-app/controller"
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CheckAuthentication(c *fiber.Ctx) bool {
	sess, _ := controller.Store.Get(c)
	username := sess.Get("login_success")
	log.Println("Session username:", username) // Log the username from the session
	return username != nil
}

func IsAuthenticated(c *fiber.Ctx) error {
	log.Println("Checking authentication...") // Log when the authentication check starts

	if !CheckAuthentication(c) {
		log.Println("User not authenticated, redirecting to login") // Log when redirecting
		return c.Redirect("/login")
	}

	log.Println("User authenticated, proceeding to the next middleware") // Log when proceeding
	return c.Next()
}

func CheckSession(c *fiber.Ctx) error {
	log.Println("Checking session...") // Log when the session check starts
	userLogin := controller.GetSessionUser(c)
	sess, _ := controller.Store.Get(c)

	if userLogin.SessionID != sess.Get("sessionID") {
		log.Println("Session mismatch, redirecting to logout") // Log when session mismatch
		return c.Redirect("/logout")
	}

	log.Println("Session valid, proceeding to the next middleware") // Log when session is valid
	return c.Next()
}

// Hàm kiểm tra phần tử có nằm trong slice hay không
func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func CheckPermission(c *fiber.Ctx) error {
	requiredPermission := c.Route().Name // hoặc sử dụng c.Route().Path nếu bạn muốn dựa trên URL

	userLogin := controller.GetSessionUser(c)

	roleId := userLogin.RoleID

	var rolePermission []entity.RolePermission
	var permission entity.Permission

	if err := database.DB.Where("role_id = ?", roleId).Find(&rolePermission).Error; err != nil {
		log.Println(err)
	}

	if err := database.DB.Where("permission = ?", requiredPermission).Find(&permission).Error; err != nil {
		log.Println(err)
	}

	var permissionID []int

	for _, s := range rolePermission {
		permissionID = append(permissionID, s.PermissionID)
	}

	log.Println("check permission", permission, rolePermission)
	log.Println("mang cac permission tuong ung voi role nguoi dung :", permissionID)
	log.Println("role", roleId, requiredPermission)

	if contains(permissionID, permission.ID) {
		log.Println("contains")
		return c.Next()
	}

	return c.Redirect("/unauthorized")
}
