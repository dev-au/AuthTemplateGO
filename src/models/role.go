package models

import "time"

type Role struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Users     []User    `gorm:"foreignKey:RoleID" json:"users"`
}
