package structureDAO

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleID uuid.UUID `gorm:"uniqueIndex"`
	Name   string
}
