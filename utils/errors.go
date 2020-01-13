package utils

import "errors"

var (
	// 短信相关
	SMSMaxLimitNumErr = errors.New("vcode count more than limit")

	// 系统错误相关
	SysErr = errors.New("System busy, please refresh later. ")

	// 请求参数验证错误
	ValidPhoneErr = errors.New("The length of mobile phone number is between 6-13 characters, for example: 12345678")
	ValidVCodeTypeErr = errors.New("VCode type invalid. ")
)

var (
	// 短信错误码
	ErrSMSMaxLimitNum = "SMS0001"

	// 系统错误码
	//ErrSys = "SYS0001"
)

var Errs = map[error]string{
	SMSMaxLimitNumErr: ErrSMSMaxLimitNum,
	//SysErr:ErrSys,
}

func ParseErr(e error) string {
	return Errs[e]
}
