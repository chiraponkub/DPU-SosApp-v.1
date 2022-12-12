package structure

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Email       *string
	PhoneNumber string
	Password    string
	FirstName   string
	LastName    string
	Birthday    time.Time
	Gender      string
	IDCard      string
	Photo       *string
	RoleID      uint
	AddressID   *uint
}

type LogLogin struct {
	gorm.Model
	UserID uint
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
