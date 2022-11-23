package structureDAO

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	UserID      uuid.UUID `gorm:"uniqueIndex"`
	Email       *string
	PhoneNumber string
	Password    string
	FirstName   string
	LastName    string
	Birthday    string
	Gender      string
	IDCard      string
	Photo       *string
	RoleID      uuid.UUID
}

type LogLogin struct {
	gorm.Model
	UserID uuid.UUID
	System string
	IP     string
}

type OTP struct {
	gorm.Model
	PhoneNumber string
	Key         int
	VerifyCode  string
	Expired     time.Time
	Active      bool
}
