package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string     `gorm:"not null" json:"title"`
	Description string     `gorm:"not null" json:"description"`
	State       string     `gorm:"not null" json:"state"`
	GroupID     uint       `gorm:"not null" json:"group_id"`
	Group       Group      `gorm:"foreignKey:GroupID" json:"group"`
	TaskUser    []TaskUser `gorm:"foreignKey:TaskID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"task_user"`
}
