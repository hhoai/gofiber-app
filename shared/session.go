package shared

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func InitSession() {
	Store = session.New()
}
