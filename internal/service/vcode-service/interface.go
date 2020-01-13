package vcode_service


type VCodeContext interface {
	GetVCodeType() string
	GetPhone() string
	SetResult(data interface{})
	GetResult() interface{}
}