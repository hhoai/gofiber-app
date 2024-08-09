package migration

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.UserEntity{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
