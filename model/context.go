package model

type VCodeContext struct {
	Param     interface{} `json:"param"`
	VCodeType string      `json:"v_code_type"`
	Phone     string      `json:"phone"`
	Result    interface{} `json:"result"`
}

func (c *VCodeContext) GetParam() interface{} {
	panic("implement me")
}

func (c *VCodeContext) GetVCodeType() string {
	return c.VCodeType
}

func (c *VCodeContext) GetPhone() string {
	return c.Phone
}

func (c *VCodeContext) SetResult(data interface{}) {
	c.Result = data
}

func (c *VCodeContext) GetResult() interface{} {
	return c.Result
}
