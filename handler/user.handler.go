package handler

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []entity.UserEntity
	// Truy vấn tất cả các bản ghi
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	log.Printf("All users: %+v", users)

	return ctx.JSON(fiber.Map{
		"user": users,
	})

	// return nil
}

// func UserHandlerCreate(ctx *fiber.Ctx) error {
// 	const MYSQL = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := MYSQL

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("can't connect database")
// 	}

// 	// Tạo bản ghi mới
// 	user := entity.UserEntity{ID: 3, Name: "John Doe", Password: "john", Email: "johndoe@example.com", Address: "address", Phone: "012345789", Role: 1}
// 	result := db.Create(&user)
// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"data": "user",
// 	})

// 	// return nil
// }
