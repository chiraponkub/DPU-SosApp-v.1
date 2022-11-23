package restapi

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/constant"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singup/request"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/response"
	"github.com/labstack/echo/v4"
)

func (ctrl Controller) SendOTP(c echo.Context) error {
	var request = new(request.PhoneNumber)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, err)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}
	resp, err := ctrl.SentOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) VerifyOTP(c echo.Context) error {
	var request = new(request.OTP)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	err = ctrl.VerifyOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "VerifySuccess"
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) CreateUser(c echo.Context) error {
	var request = new(request.Account)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, err)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	err = ctrl.CreateUserLogin(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "Succeed"
	return response.EchoSucceed(c, res)

}

func (ctrl Controller) SignInUser(c echo.Context) error {

	return response.EchoSucceed(c, "")
}

func (ctrl Controller) SignUpAdmin(c echo.Context) error {

	return response.EchoSucceed(c, "")
}
