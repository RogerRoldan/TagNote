package models

type Group struct {
	ID                 uint   `gorm:"primaryKey" json:"id"`
	Title              string `gorm:"not null" json:"title"`
	Description        string `gorm:"not null" json:"description"`
	GroupsUser         []User `gorm:"many2many:group_user" json:"groups_user"`
	AdministratorGroup []User `gorm:"many2many:administrator_group" json:"administrator_group"`
}
