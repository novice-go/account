package model


type Account struct {
	Base
	Phone string `gorm:"index" json:"phone"`
}