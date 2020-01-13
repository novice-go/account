package service

import (
	"wdkj/account/internal/service/account-service"
	vcode_service "wdkj/account/internal/service/vcode-service"
)

type AccountService interface {
	Login(c account_service.AccContext) error
	Register(c account_service.AccContext) error
}

type VCodeService interface {
	GenVCode(c vcode_service.VCodeContext) error
}

type LoginManager struct {
	accountService AccountService
	vcodeService VCodeService
}

func NewLoginManager(accountService AccountService, vcodeService VCodeService) *LoginManager {
	return &LoginManager{accountService: accountService, vcodeService: vcodeService}
}

func (m *LoginManager) Login(c account_service.AccContext) error {
	return m.accountService.Login(c)
}

func (m *LoginManager) Register(c account_service.AccContext) error {
	return m.Register(c)
}

func (m *LoginManager) GenVCode(c vcode_service.VCodeContext) error {
	return m.vcodeService.GenVCode(c)
}
