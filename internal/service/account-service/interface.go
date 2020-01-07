package account_service


type AccContext interface {
	GetPhone() string
	GetParam() interface{}
	SetResult(data interface{})
	GetResult() interface{}
}