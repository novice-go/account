package account_service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"wdkj/account/model"
)

type AccountDAO interface {
	QueryAccountByPhone(phone string) (*model.Account, error)
	CreateAccount (account *model.Account) error
}

type TokenDAO interface {
	GenToken(phone string) (token string, err error)
	AuthToken(token string) error
}

type AccountService struct {
	dao AccountDAO
}

func NewAccountService(dao AccountDAO) *AccountService {
	return &AccountService{dao: dao}
}

func (a *AccountService) Register(c AccContext) error {
	data, ok := c.GetParam().(*model.Account)
	if !ok {
		return errors.New("##AccountService.Register param to model.account fail")
	}

	return a.dao.CreateAccount(data)
}

func (a *AccountService) Login(c AccContext) error {
	// 验证用户是否存在
	resp, err := a.dao.QueryAccountByPhone(c.GetPhone())
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// TODO 生成token
	c.SetResult(resp)
	return nil
}



