package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the reveue : ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate : ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt - ebt*taxRate/100
	ratio := ebt / profit

	fmt.Println("EBT : ", ebt)
	fmt.Println("Profit : ", profit)
	fmt.Println("Ratio : ", ratio)

}
