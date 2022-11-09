package restapi

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model"
	rdbmsstructure "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/structure"
	"strings"
)

func (ctrl Controller) AddRoleCon(req *model.AddRole) (Error error) {
	var newReq rdbmsstructure.Role
	newReq.Name = strings.ToLower(req.Name)

	res, err := ctrl.Access.RDBMS.GetRoleDB(newReq)
	if res.Name == req.Name {
		Error = errors.New("มี Role นี้ในระบบแล้ว")
		return
	}

	role := rdbmsstructure.Role{
		Name: req.Name,
	}
	err = ctrl.Access.RDBMS.AddRoleDB(role)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl Controller) GetRoleListCon() (response []model.GetRoleList, Error error) {
	data, err := ctrl.Access.RDBMS.GetRoleListDB()
	if err != nil {
		Error = err
		return
	}
	for _, m1 := range data {
		arr := model.GetRoleList{
			Name: m1.Name,
		}
		response = append(response, arr)
	}
	return
}
