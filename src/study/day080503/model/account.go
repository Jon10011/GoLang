package model

import "fmt"

type account struct {
	accountNum string
	accountPwd string
	money      float64
}

func NewAccount(a string) *account {
	if len(a) >= 6 && len(a) <= 10 {
		return &account{
			accountNum: a,
		}
	} else {
		fmt.Println("输入账号错误")
		return nil
	}
}

func (account *account) SetPwd(password string) {
	if len(password) == 6 {
		account.accountPwd = password
	} else {
		fmt.Println("输入密码格式错误，请输入6为密码")
	}
}

func (account *account) SetMoney(money float64) {
	if money >= 20.0 {
		account.money = money
	} else {
		fmt.Println("余额太少啦")
	}
}
