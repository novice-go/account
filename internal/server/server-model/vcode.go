package server_model

import "wdkj/account/utils"

type GenVCodeReq struct {
	VCodeType string `json:"v_code_type"`
	Phone     string `json:"phone"`
}

func (req *GenVCodeReq) IsValid() error {
	switch {
	case len(req.Phone) < 6 || len(req.Phone) > 13: // 根据实际验证咯
		return utils.ValidPhoneErr
	case req.VCodeType != utils.VCodeTypeLogin && req.VCodeType != utils.VcodeTypeRegister:
		return utils.ValidVCodeTypeErr
	default:
		return nil
	}
}
