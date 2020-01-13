package vcode_service

import (
	"go.uber.org/zap"
	"time"
	"wdkj/account/model"
	"wdkj/account/utils"
	"wdkj/account/utils/log"
)

type VCodeFLowDAO interface {
	GetVCodeCountByPhone(phone, typ, status string) (count uint, err error)
	SaveVCodeFlow(flow *model.SMSFlow) error
}

type VCodeDAO interface {
	SaveVCode(flow *model.VCode) error
}

type Sender interface {
	SendSMS(vcode, typ, phone string) (flow *model.SMSFlow, err error)
}

type VCodeService struct {
	flowDAO VCodeFLowDAO
	dao     VCodeDAO
	sender  Sender
}

func NewVCodeService(flowDAO VCodeFLowDAO, dao VCodeDAO, sender Sender) *VCodeService {
	return &VCodeService{flowDAO: flowDAO, dao: dao, sender: sender}
}

func (s *VCodeService) GenVCode(c VCodeContext) error {
	// 校验该手机当天该类型验证码发送次数
	count, err := s.flowDAO.GetVCodeCountByPhone(c.GetPhone(), c.GetVCodeType(), utils.StatusSuccess)
	if err != nil {
		return err
	}
	if count >= utils.SMSMaxLimitNum {
		return utils.SMSMaxLimitNumErr
	}

	// 发送验证码, 返回发送流水
	flow, err := s.sender.SendSMS(utils.GenNumber(), c.GetVCodeType(), c.GetPhone())
	if err != nil {
		log.Logger.Error("send sms fail",
			zap.String("VCodeType", c.GetVCodeType()),
			zap.String("phone", c.GetPhone()),
			zap.Error(err),
		)
	}

	// 保存流水
	if err = s.flowDAO.SaveVCodeFlow(flow); err != nil {
		return err
	}

	// 保存验证码信息
	if err := s.dao.SaveVCode(&model.VCode{
		Account:   model.Account{Phone:c.GetPhone()},
		Used:      false,
		VCode:     flow.Content,
		ErrTimes:  0,
		VCodeType: c.GetVCodeType(),
	}); err != nil {
		return err
	}

	c.SetResult(count+1)
	return nil
}

// 暂不开发短信发送功能, 暂用mock
type mockSenderImpl struct{}

func NewMockSenderImpl() *mockSenderImpl {
	return &mockSenderImpl{}
}

func (m *mockSenderImpl) SendSMS(vcode, typ, phone string) (*model.SMSFlow, error) {
	t := time.Now()
	flow := &model.SMSFlow{
		Account:    model.Account{Phone:phone},
		Content:    vcode,
		FlowNo:     utils.GetMsgId(),
		SendStatus: utils.StatusSuccess,
		SendType:   typ,
		SendAt:     &t,
	}

	// send sms
	return flow, nil
}


