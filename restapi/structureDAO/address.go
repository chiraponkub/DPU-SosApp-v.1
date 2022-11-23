package structureDAO

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Address     string
	SubDistrict string
	District    string
	Province    string
	PostalCode  string
	Country     string
	UserID      uuid.UUID
}
