package model

type ResponseMain struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	OTP  OTP
}

type OTP struct {
	OTP        int    `json:"otp"`
	VerifyCode string `json:"verifyCode"`
}
