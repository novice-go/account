package ctx


type BaseContext interface {
	GetParam() interface{}
	SetResult(data interface{})
	GetResult() interface{}
}

type AccContext interface {
	VCodeContext
	GetVCode() string
}

type VCodeContext interface {
	BaseContext
	GetVCodeType() string
	GetPhone() string
}
