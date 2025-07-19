package models

import (
	"gorm.io/gorm"
)

type Role struct {
	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Users []User `gorm:"foreignKey:RoleID" json:"users"`
	gorm.Model
}
