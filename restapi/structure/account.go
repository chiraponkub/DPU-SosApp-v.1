package structure

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email       string
	Username    string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
	Birthday    string
	Sex         string
	IDCard      string
	Photo       string
	Position    string
}
