package restapi

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/structureDAO"
	"gorm.io/gorm"
)

func (factory GORMFactory) GetAccountDB(req structureDAO.Account) (response *structureDAO.Account, Error error) {
	var data = new(structureDAO.Account)
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
