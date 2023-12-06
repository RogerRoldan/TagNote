package models

type Group struct {
	ID                 uint                 `gorm:"primaryKey"`
	Title              string               `gorm:"not null"`
	Description        string               `gorm:"not null"`
	GroupsUser         []GroupUser          `gorm:"foreignKey:ID"`
	AdministratorGroup []AdministratorGroup `gorm:"foreignKey:ID"`
}
