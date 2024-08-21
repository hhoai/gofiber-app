package middleware

import (
	"fiber-app/controller"
	"fiber-app/shared"
	"log"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	log.Println("check isAuthenticated middleware")
	sess, _ := shared.Store.Get(c)
	var username = sess.Get("username")
	if username == nil {
		return c.Redirect("/login")
	}
	return c.Next()
}

func CheckSession(c *fiber.Ctx) error {
	log.Println("check session middleware")
	userLoginNow := controller.GetUserSessionNow(c)
	sess, _ := shared.Store.Get(c)
	log.Println(sess.Get("sessionid"))
	if userLoginNow.Sessionid != sess.Get("sessionid") {
		return c.Redirect("/logout")
	}
	return c.Next()
}
