package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetSalesData(c *fiber.Ctx) error {
	var sales []entity.SalesData
	database.DB.Find(&sales)

	var product1 []entity.SalesData
	database.DB.Where("product_id = 1").Find(&product1)

	var product2 []entity.SalesData
	database.DB.Where("product_id = 2").Find(&product2)

	var account []entity.UserWithRowNumber
	result := database.DB.Table("user_entities").
		Joins("INNER JOIN roles ON user_entities.role_id = roles.id").
		Select("ROW_NUMBER() OVER (ORDER BY id) AS RowNumber, user_entities.id, name, email, address, phone, roles.role AS role_name").
		Find(&account)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.JSON(
		fiber.Map{
			"data":      sales,
			"product_1": product1,
			"product_2": product2,
			"account":   account,
		},
	)
}

func SalesController(c *fiber.Ctx) error {
	return c.Render("sales", fiber.Map{
		// "SaleData": sales,
		"Ctx": c,
	}, "layouts/main")
}
