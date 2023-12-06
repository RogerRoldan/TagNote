package models

type Group struct {
	ID                 uint                 `gorm:"primaryKey" json:"id"`
	Title              string               `gorm:"not null" json:"title"`
	Description        string               `gorm:"not null" json:"description"`
	GroupsUser         []GroupUser          `gorm:"foreignKey:ID" json:"groupsUser_id"`
	AdministratorGroup []AdministratorGroup `gorm:"foreignKey:ID" json:"administratorGroup_id"`
}
