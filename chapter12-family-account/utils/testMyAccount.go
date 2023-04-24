package utils

import "fmt"

type FamilyAccount struct {
	balance        float64 // 账户余额
	accountDetails []AccountDetail
}

type AccountDetail struct {
	operation string  // 收入还是支出
	money     float64 // 操作金额
	note      string  // 收支说明

}

// 工厂模式
func NewFamilyAccount(balance float64) *FamilyAccount {
	return &FamilyAccount{
		balance: balance,
	}
}

func NewAccountDetail(operation string, money float64, note string) *AccountDetail {
	return &AccountDetail{
		operation: operation,
		money:     money,
		note:      note,
	}
}

func (fa *FamilyAccount) ShowMenu() {
	fmt.Println("------------- 家庭收支记账软件 -------------")
	fmt.Println("                1 收支明细                  ")
	fmt.Println("                2 登记收入                  ")
	fmt.Println("                3 登记支出                  ")
	fmt.Println("                4 退    出                  ")
	fmt.Println()
	fmt.Print("                请选择(1-4): ")
}

func (fa *FamilyAccount) ShowDetails() {
	// fmt.Println("显示详情")
	fmt.Println()
	fmt.Println("----当前收支明细记录----")
	fmt.Printf("收支\t账户金额\t收支金额\t说明\n")

	for _, ad := range fa.accountDetails {
		fmt.Printf("%s\t%f\t%f\t%s\n", ad.operation, fa.balance, ad.money, ad.note)
	}
	fmt.Println()
}

func (fa *FamilyAccount) In() {
	// fmt.Println("登记收入")

	var money float64
	var note string
	fmt.Print("本次收入金额：")
	fmt.Scanln(&money)
	fmt.Print("本次收入说明：")
	fmt.Scanln(&note)

	fa.balance += money

	ad := NewAccountDetail("收入", money, note)
	fa.accountDetails = append(fa.accountDetails, *ad)
}

func (fa *FamilyAccount) Out() {
	// fmt.Println("登记支出")

	var money float64
	var note string
	fmt.Print("本次支出金额：")
	fmt.Scanln(&money)
	fmt.Print("本次支出说明：")
	fmt.Scanln(&note)

	if money <= fa.balance {
		fa.balance -= money
		ad := NewAccountDetail("支出", money, note)
		fa.accountDetails = append(fa.accountDetails, *ad)
	} else {
		fmt.Println("支出金额大于账户余额，支出失败")
	}
}

func (fa *FamilyAccount) Exit() {
	fmt.Println("程序退出")
}

func (fa *FamilyAccount) Err() {
	fmt.Println("输入有误，重新输入")
}

func (fa *FamilyAccount) Choose(num int) {
	switch num {
	case 1:
		fa.ShowDetails()
	case 2:
		fa.In()
	case 3:
		fa.Out()
	case 4:
		fa.Exit()
	default:
		fa.Err()
	}

}
