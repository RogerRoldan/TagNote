package models

import "gorm.io/gorm"

type TaskUser struct {
	gorm.Model
	TaskID uint `gorm:"not null" json:"task_id"`
	Task   Task `gorm:"foreignKey:TaskID" json:"task"`
	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`
}
