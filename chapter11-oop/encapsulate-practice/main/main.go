package main

import (
	"fmt"
	account "go-practice/chapter11-oop/encapsulate-practice/model"
)

func main() {
	account := account.NewAccount("ananan", "123123", 21)
	fmt.Println(account)
	if account != nil {
		fmt.Println("创建账户成功")
	} else {
		fmt.Println("创建账户失败")
	}

	account.SetAccountId("ningning")
	fmt.Println(account)

	account.SetPwd("1234567")
	fmt.Println(account)

	account.SetBalance(1234)
	fmt.Println(account)
	fmt.Println()

	fmt.Println("账户:", account.GetAccountId(), "密码:", account.GetPwd(), "余额:", account.GetBalance())

}
