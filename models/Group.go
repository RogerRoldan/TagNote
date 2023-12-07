package models

type Group struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title              string `gorm:"not null" json:"title"`
	Description        string `gorm:"not null" json:"description"`
	Task               []Task `gorm:"foreignKey:ID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"task"`
	GroupsUser         []User `gorm:"many2many:group_user" json:"groups_user"`
	AdministratorGroup []User `gorm:"many2many:administrator_group" json:"administrator_group"`
}
