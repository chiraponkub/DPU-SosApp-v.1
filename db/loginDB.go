package db

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/db/structure"
	"gorm.io/gorm"
)

func (factory GORMFactory) GetAccountDB(req structure.Account) (response *structure.Account, Error error) {
	var data = new(structure.Account)
	err := factory.client.Where("phone_number = ?", req.PhoneNumber).Find(&data).Error
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
