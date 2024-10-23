package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique;not null" json:"email" validate:"email,required"`
	Password        string `gorm:"not null" json:"password,omitempty" validate:"required"`
	Name            string `gorm:"not null" json:"name" validate:"required"`
	ProfilePic      string `gorm:"not null" json:"profile_pic"`
	LoginLengthTime uint32 `json:"-" validate:"number"`
	// server data
	Validated      bool      `gorm:"not null;default:false" json:"validated" validate:"boolean"`
	ValidationCode string    `gorm:"not null" json:"validation_code,omitempty"`
	Admin          bool      `gorm:"not null;default:false" json:"-"`
	SuperAdmin     bool      `gorm:"not null;default:false" json:"-"`
	Code           string    `json:"-"`
	CodeExpire     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}
