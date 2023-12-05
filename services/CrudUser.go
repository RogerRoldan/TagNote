package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func GetUsers(db *gorm.DB) ([]models.User, bool) {
	var users []models.User
	usersList := db.Find(&users)
	if usersList.Error != nil {
		log.Println("Error al obtener los usuarios")
		return []models.User{}, false
	}
	
	log.Println("Usuarios encontrados exitosamente...")
	return users, true
}

func CreateUser(db *gorm.DB, user models.User) (string, bool) {
	userCreate := db.Create(&user)
	if userCreate.Error != nil {
		log.Println("Error al crear el usuario")
		return "Error al crear el usuario", false
	}
	log.Println("Usuario creado exitosamente...")
	return "Usuario creado exitosamente...", true
}

func GetUserId(db *gorm.DB, id int) (models.User, bool) {
	var user models.User
	response := db.First(&user, id)
	if response.Error != nil {
		log.Println("Usuario no encontrado...")
		return models.User{}, false
	}

	log.Println("Usuario encontrado exitosamente...")
	return user, true
}

func DeleteUser(db *gorm.DB, id int) (string, bool) {
	var user models.User
	db.First(&user, id)
	if user == (models.User{}) {
		log.Println("Usuario no encontrado...")
		return "Usuario no encontrado...", false
	}

	userDelete := db.Delete(&user)
	if userDelete.Error != nil {
		log.Println("Error al eliminar el usuario")
		return "Error al eliminar el usuario", false
	}
	log.Println("Usuario eliminado exitosamente...")
	return "Usuario eliminado exitosamente...", true
}

func UpdateUser(db *gorm.DB, user models.User) (string, bool) {
	userM := db.Save(&user)
	if userM.Error != nil {
		log.Println("Error al actualizar el usuario")
		return "Error al actualizar el usuario", false
	}

	log.Println("Usuario actualizado exitosamente...")
	return "Usuario actualizado exitosamente...", true
}
