package main

import "fmt"

// getting user inputs from terminal
func main() {
	var employeeId string
	var salary int
	var age int

	fmt.Print("Enter your Employee Id: ")
	fmt.Scan(&employeeId)

	fmt.Print("Enter your salary: ")
	fmt.Scan(&salary)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("EmployeeId: ", employeeId, " \nSalary: ", salary, " \nAge: ", age)
}
