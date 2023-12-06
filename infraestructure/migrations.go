package infraestructure

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
)

func Migrate() {
	var db *gorm.DB = GetConnection()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Task{})
}
