package infraestructure

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
)

func Migrate() {
	var db *gorm.DB = GetConnection()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.GroupUser{})
	db.AutoMigrate(&models.AdministratorGroup{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.TaskAssigned{})
}
