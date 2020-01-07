package service

import "wdkj/account/internal/service/account-service"

type AccountService interface {
	Login(c account_service.AccContext) error
	Register(c account_service.AccContext) error
}

type LoginManager struct {
	accountService AccountService
}

func(m *LoginManager) Login (c account_service.AccContext) error {
	return m.accountService.Login(c)
}

func (m *LoginManager) Register (c account_service.AccContext) error {
	return m.Register(c)
}