package models

import "gorm.io/gorm"

type GroupUser struct {
	gorm.Model
	GroupID uint  `gorm:"not null" json:"group_id"`
	Group   Group `gorm:"foreignKey:GroupID" json:"group"`
	UserID  uint  `gorm:"not null" json:"user_id"`
	User    User  `gorm:"foreignKey:UserID" json:"user"`
}
