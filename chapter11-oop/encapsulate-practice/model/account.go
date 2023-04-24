package model

import "fmt"

type account struct {
	accountId string
	pwd       string
	balance   float64
}

func NewAccount(accountId string, pwd string, balance float64) *account {
	if len(accountId) < 6 || len(accountId) > 10 {
		fmt.Println("账户长度有误")
		return nil
	}

	if balance <= 20 {
		fmt.Println("余额有误")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码有误")
		return nil
	}

	return &account{
		accountId: accountId,
		pwd:       pwd,
		balance:   balance,
	}
}

func (account *account) SetAccountId(accountId string) {
	account.accountId = accountId
}

func (account *account) GetAccountId() string {
	return account.accountId
}

func (account *account) SetPwd(pwd string) {
	account.pwd = pwd
}

func (account *account) GetPwd() string {
	return account.pwd
}

func (account *account) SetBalance(balance float64) {
	account.balance = balance
}

func (account *account) GetBalance() float64 {
	return account.balance
}
