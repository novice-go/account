package service

import (
	"wdkj/account/internal/ctx"
)

type AccountService interface {
	Login(c ctx.AccContext) error
	Register(c ctx.AccContext) error
}

type VCodeService interface {
	GenVCode(c ctx.VCodeContext) error
	AuthVCode(phone, vcode, typ string) error
}

type LoginManager struct {
	accountService AccountService
	vcodeService VCodeService
}

func NewLoginManager(accountService AccountService, vcodeService VCodeService) *LoginManager {
	return &LoginManager{accountService: accountService, vcodeService: vcodeService}
}

func (m *LoginManager) Login(c ctx.AccContext) error {
	return m.accountService.Login(c)
}

func (m *LoginManager) Register(c ctx.AccContext) error {
	if err := m.vcodeService.AuthVCode(c.GetPhone(), c.GetVCode(), c.GetVCodeType()); err != nil {
		return err
	}

	return m.accountService.Register(c)
}

func (m *LoginManager) GenVCode(c ctx.VCodeContext) error {
	return m.vcodeService.GenVCode(c)
}
