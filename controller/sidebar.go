package controller

import (
	"fiber-app/database"
	"fiber-app/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Sidebar(c *fiber.Ctx) []entity.SidebarItem {
	user := GetSessionUser(c)

	var roleSidebarItem []entity.RoleSidebarItem
	if err := database.DB.Where("role_id = ?", user.RoleID).Find(&roleSidebarItem); err != nil {
		log.Println(err)
	}

	var sidebarItems []entity.SidebarItem
	for _, item := range roleSidebarItem {
		var sidebarItem entity.SidebarItem
		if err := database.DB.Where("item_id = ?", item.ItemID).First(&sidebarItem).Error; err != nil {
			log.Println(err)
		} else {
			sidebarItems = append(sidebarItems, sidebarItem)
		}
	}

	log.Println("sidebar: ", sidebarItems)

	return sidebarItems
}
