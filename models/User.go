package models

type User struct {
	ID                 uint                 `gorm:"primaryKey" json:"id"`
	Username           string               `gorm:"unique; not null" json:"username"`
	Password           string               `gorm:"not null" json:"password"`
	Email              string               `gorm:"unique; not null" json:"email"`
	FirstName          string               `gorm:"not null" json:"first_name"`
	LastName           string               `gorm:"not null" json:"last_name"`
	Imagen             string               `gorm:"null" json:"imagen"`
	TaskAssigned       []TaskAssigned       `gorm:"foreignKey:ID"`
	GroupsUser         []GroupUser          `gorm:"foreignKey:ID"`
	AdministratorGroup []AdministratorGroup `gorm:"foreignKey:ID"`
}
