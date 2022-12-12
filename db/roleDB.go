package db

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/db/structureDAO"
	"gorm.io/gorm"
)

func (factory GORMFactory) AddRoleDB(req structureDAO.Role) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) GetRoleListDB() (response []structureDAO.Role, Error error) {
	var data []structureDAO.Role
	err := factory.client.Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}
	response = data
	return
}

func (factory GORMFactory) GetRoleDBByName(req structureDAO.Role) (response structureDAO.Role, Error error) {
	var data structureDAO.Role
	err := factory.client.Where("name = ?", req.Name).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}
	response = data
	return
}

func (factory GORMFactory) GetRoleDBById(req structureDAO.Role) (response structureDAO.Role, Error error) {
	var data structureDAO.Role
	err := factory.client.Where("id = ?", req.ID).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}
	response = data
	return
}
