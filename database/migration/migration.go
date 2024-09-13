package migration

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.UserEntity{})
	database.DB.AutoMigrate(&entity.Role{})
	database.DB.AutoMigrate(&entity.Permission{})
	database.DB.AutoMigrate(&entity.RolePermission{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
