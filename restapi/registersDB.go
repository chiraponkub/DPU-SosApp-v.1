package restapi

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/structureDAO"
	"gorm.io/gorm"
)

func (factory GORMFactory) SendOTPDB(req structureDAO.OTP) (Error error) {
	var data []structureDAO.OTP
	db := factory.client.Model(&data).Where("phone_number = ?", req.PhoneNumber).Update("active", false).Error
	if db != nil {
		Error = db
		return
	}

	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) GetOTPDB(req structureDAO.OTP) (response *structureDAO.OTP, Error error) {
	var data = new(structureDAO.OTP)
	err := factory.client.Where("phone_number = ? and active = ?", req.PhoneNumber, true).Find(&data).Error
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

func (factory GORMFactory) CreateUserDB(req structureDAO.Account) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}
