package initializers

import (
	"Go/JWT_Go/models"
	"gorm.io/gorm"
)

func SyncDatabase(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}