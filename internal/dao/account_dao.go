package dao

import "wdkj/account/model"

type AccountDB interface {

}

type AccountDAO struct {

}

func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

func (a AccountDAO) QueryAccountByPhone(phone string) (*model.Account, error) {
	panic("implement me")
}

func (a AccountDAO) CreateAccount(account *model.Account) error {
	panic("implement me")
}



