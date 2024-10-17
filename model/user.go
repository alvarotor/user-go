package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string    `gorm:"uniqueIndex;not null" json:"email"`
	Password        string    `gorm:"not null" json:"password,omitempty"`
	Name            string    `gorm:"not null" json:"name"`
	Age             uint      `gorm:"not null;default:0" json:"age"`
	Gender          string    `gorm:"not null" json:"gender"`
	CountryOrigin   string    `json:"country_origin"`
	ProfilePic      string    `gorm:"not null" json:"profile_pic"`
	Validated       bool      `gorm:"not null;default:false" json:"validated"`
	ValidationCode  string    `gorm:"not null" json:"validation_code,omitempty"`
	Admin           bool      `json:"-"`
	SuperAdmin      bool      `json:"-"`
	LoginLengthTime uint      `json:"-"`
	Code            string    `json:"-"`
	CodeExpire      time.Time `json:"-"`
}
