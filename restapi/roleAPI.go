package restapi

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/response"
	"github.com/labstack/echo/v4"
)

func (ctrl Controller) AddRole(c echo.Context) error {
	var request = new(model.AddRole)
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, "Succeed")
	}

	err = ValidateStruct(request)
	if err != nil {
		return response.EchoError(c, 400, err.Error())
	}
	err = ctrl.AddRoleCon(request)
	if err != nil {
		return response.EchoError(c, 400, err.Error())
	}
	return response.EchoSucceed(c, "Succeed")
}

func (ctrl Controller) GetRoleList(c echo.Context) error {
	responses, err := ctrl.GetRoleListCon()
	if err != nil {
		return response.EchoError(c, 400, err.Error())
	}
	return response.EchoSucceed(c, responses)
}
