package structure

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	Username         string `gorm:"unique"`
	Password         string
	Firstname        string
	Lastname         string
	Email            *string
	PhoneNumber      string `gorm:"size:10;unique"`
	Birthday         time.Time
	Gender           string
	IDCard           int
	PathImageProfile *string
	DeletedBy        *int
	Workplace        *string
	AddressID        *uint
	RoleID           *uint
}

type Address struct {
	gorm.Model
	Address     string
	SubDistrict string
	District    string
	Province    string
	PostalCode  string
	Country     string
	DeletedBy   uint
}

type OTP struct {
	gorm.Model
	PhoneNumber string
	Key         int
	VerifyCode  string
	Expired     time.Time
	Active      bool
}

type LogLogin struct {
	gorm.Model
	UserID uint
	System string
	IP     string
}

type IDCard struct {
	gorm.Model
	IDCardText string
	PathImage  string
	DeletedBy  uint
}
