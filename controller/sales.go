package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetSalesData(c *fiber.Ctx) error {
	var sales []entity.SalesData
	database.DB.Find(&sales)

	var product1 []entity.SalesData
	database.DB.Where("product_id = 1").Find(&product1)

	var product2 []entity.SalesData
	database.DB.Where("product_id = 2").Find(&product2)

	draw, _ := strconv.Atoi(c.Query("draw"))
	start, _ := strconv.Atoi(c.Query("start"))
	length, _ := strconv.Atoi(c.Query("length"))
	searchValue := c.Query("search[value]")

	var totalRecords int64
	var filteredRecords int64
	var users []entity.UserEntity

	// Get total number of records
	database.DB.Table("user_entities").Count(&totalRecords)

	// Apply search filter if provided
	query := database.DB.Table("user_entities")
	if searchValue != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+searchValue+"%", "%"+searchValue+"%")
	}

	// Get total number of filtered records
	query.Count(&filteredRecords)

	// Apply pagination
	// database.DB.Table("user_entities").Offset(start).Limit(length).Find(&users)
	database.DB.Offset(start).Limit(length).Find(&users)

	var account []entity.UserWithRowNumber
	result := query.
		Joins("INNER JOIN roles ON user_entities.role_id = roles.id").
		Select("ROW_NUMBER() OVER (ORDER BY id) AS RowNumber, user_entities.id, name, email, address, phone, roles.role AS role_name").
		Offset(start).
		Limit(length).
		Find(&account)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	// Prepare response
	response := map[string]interface{}{
		"draw":            draw,
		"recordsTotal":    totalRecords,
		"recordsFiltered": filteredRecords,
		"data":            account,
		// "data":      sales,
		"product_1": product1,
		"product_2": product2,
	}

	// Return JSON response
	return c.JSON(response)
}

func SalesController(c *fiber.Ctx) error {
	return c.Render("sales", fiber.Map{
		// "SaleData": sales,
		"Ctx": c,
	}, "layouts/main")
}
