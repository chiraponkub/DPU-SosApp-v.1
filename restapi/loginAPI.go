package restapi

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/constant"
	singin "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singin/request"
	singinResp "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singin/response"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/response"
	"github.com/labstack/echo/v4"
)

func (ctrl Controller) SignInUser(c echo.Context) error {
	var request = new(singin.Login)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, err)
	}
	token, err := ctrl.LoginLogic(request)
	if err != nil {
		res.Msg = err.Error()
		res.Code = constant.ErrorCode
		return response.EchoSucceed(c, res)
	}

	resp := singinResp.TokenRes{
		Token: token,
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, resp)
}

func (ctrl Controller) SignUpAdmin(c echo.Context) error {

	return response.EchoSucceed(c, "")
}
