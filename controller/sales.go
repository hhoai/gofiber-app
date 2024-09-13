package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"

	"github.com/gofiber/fiber/v2"
)

func GetSalesData(c *fiber.Ctx) error {
	var sales []entity.SalesData
	database.DB.Find(&sales)
	return c.JSON(sales)
}

func SalesController(c *fiber.Ctx) error {
	return c.Render("sales", fiber.Map{
		// "SaleData": sales,
		"Ctx": c,
	}, "layouts/main")
}
