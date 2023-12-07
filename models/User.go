package models

type User struct {
	ID                 uint        `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Username           string      `gorm:"unique; not null" json:"username"`
	Password           string      `gorm:"not null" json:"password"`
	Email              string      `gorm:"unique; not null" json:"email"`
	FirstName          string      `gorm:"not null" json:"first_name"`
	LastName           string      `gorm:"not null" json:"last_name"`
	Imagen             string      `gorm:"null" json:"imagen"`
	Task               []Task      `gorm:"foreignKey:UserID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"task"`
	GroupUser          []GroupUser `gorm:"foreignKey:UserID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"groups_user"`
	AdministratorGroup []Group     `gorm:"foreignKey:AdminID;gorm:constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"task"`
}
