package models

import (
	"time"

	"gorm.io/gorm"
)

type DeviceInfo struct {
	Browser                string `json:"browser"`
	BrowserVersion         string `json:"browser_version"`
	OperatingSystem        string `json:"operatingSystem"`
	OperatingSystemVersion string `json:"operatingSystem_version"`
	Cpu                    string `json:"cpu"`
	Language               string `json:"language"`
	Timezone               string `json:"timezone"`
	CookiesEnabled         bool   `json:"cookies_enabled"`
}

type User struct {
	gorm.Model
	Email           string `gorm:"uniqueIndex:idx_email;not null" validate:"email,required"`
	Password        string `gorm:"not null" validate:"required"`
	Name            string `gorm:"not null" validate:"required"`
	ProfilePic      string `gorm:"not null"`
	// server data
	Admin      bool `gorm:"not null;default:false"`
	SuperAdmin bool `gorm:"not null;default:false"`
	Validated  bool `gorm:"not null;default:false" validate:"boolean"`
	Code       string
	CodeRefresh string
	CodeExpire time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Bucket     string
	DeviceInfo
}
