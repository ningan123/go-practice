package main

import (
	"fmt"
	"go-practice/chapter12-family-account/utils"
)

func main() {
	var num int
	fa := utils.NewFamilyAccount(10000)

	for {
		fa.ShowMenu()
		fmt.Scanln(&num)
		fa.Choose(num)
		if num == 4 {
			break
		}
	}

}
