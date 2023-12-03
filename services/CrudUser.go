package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateUser(db *gorm.DB, user models.User) {
	db.Create(&user)
	log.Println("Usuario creado exitosamente...")
}

func ReadUser(db *gorm.DB, user models.User) {
	db.Find(&user)
	log.Println("Usuario encontrado exitosamente...")
}

func DeleteUser(db *gorm.DB, user models.User) {
	db.Delete(&user)
	log.Println("Usuario eliminado exitosamente...")
}

func UpdateUser(db *gorm.DB, user models.User) {
	db.Save(&user)
	log.Println("Usuario actualizado exitosamente...")
}
