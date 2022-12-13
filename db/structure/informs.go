package structure

import (
	"gorm.io/gorm"
	"time"
)

type Inform struct {
	gorm.Model
	Description         string
	PhoneNumberCallBack string
	UserID              uint
	DeletedBy           uint
	SubTypeID           uint
}

type InformImage struct {
	gorm.Model
	Image    string
	InformID uint
}

type InformNotification struct {
	InformID    uint
	UserID      uint
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	DeletedBy   uint
}
