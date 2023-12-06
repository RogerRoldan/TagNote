package models

type GroupUser struct {
	ID    uint  `gorm:"primaryKey" json:"id"`
	Group Group `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"group_id"`
	User  User  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
}
