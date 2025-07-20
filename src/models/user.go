package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	IsActive  bool      `gorm:"type:boolean;not null;default:false" json:"is_active"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	RoleID    *uint     `gorm:"null" json:"role_id"`
	Role      *Role     `gorm:"foreignKey:RoleID;constraint:OnDelete:RESTRICT" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.ID = uuid.New()
	return
}
