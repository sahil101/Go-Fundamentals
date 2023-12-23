package main

import "fmt"

func main() {

	var Salary int = 100000

	var salaryPointer *int

	salaryPointer = &Salary

	fmt.Print("My Salary: ")
	fmt.Println(*salaryPointer, salaryPointer)
}

// output: My Salary: 100000 0xc000014080

// prevent copies of same data
// data mutation
