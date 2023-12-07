package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Title       string      `gorm:"not null" json:"title"`
	Description string      `gorm:"not null" json:"description"`
	Task        []Task      `gorm:"foreignKey:GroupID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"task"`
	GroupUser   []GroupUser `gorm:"foreignKey:GroupID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"groups_user"`
	AdminID     uint        `gorm:"not null" json:"admin_id"`
	Admin       User        `gorm:"foreignKey:AdminID" json:"admin"`
}
