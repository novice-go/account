package model

type Account struct {
	Base
	Phone string `gorm:"index" json:"phone" name:"手机号码"` // 索引
}

type VCode struct {
	Account
	Used      bool   `json:"used" name:"是否已被使用"`
	VCode     string `json:"v_code" name:"验证码"`
	ErrTimes  int    `json:"err_times" name:"错误次数"`
	VCodeType string `json:"v_code_type" name:"验证码类型"`
}
