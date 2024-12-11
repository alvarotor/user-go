package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string `gorm:"uniqueIndex:idx_email;not null" validate:"email,required"`
	Password        string `gorm:"not null" validate:"required"`
	Name            string `gorm:"not null" validate:"required"`
	ProfilePic      string `gorm:"not null"`
	LoginLengthTime uint32 `validate:"number"`
	// server data
	Admin      bool `gorm:"not null;default:false"`
	SuperAdmin bool `gorm:"not null;default:false"`
	Validated  bool `gorm:"not null;default:false" validate:"boolean"`
	Code       string
	CodeExpire time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Bucket     string
}
