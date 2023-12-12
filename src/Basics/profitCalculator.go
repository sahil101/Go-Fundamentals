package main

import "fmt"

func main() {
	var revenue, expenses int
	var taxRate float64
	
	fmt.Print("Enter revenue of your company: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses of your company: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	fmt.Println("Earnings Before Tax: ", EBT)

	var EAT float64 = float64(EBT) - ((taxRate / 100.0) * float64(EBT))
	fmt.Println("Profits: ", EAT)

	var ratio float64 = float64(EBT) / EAT
	fmt.Println("EBT / profits: ", ratio)

}