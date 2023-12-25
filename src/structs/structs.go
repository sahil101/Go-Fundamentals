package main

import (
	"errors"
	"fmt"
)

type user struct {
	firstName  string
	lastName   string
	age        int
	employeeId int
	location   string
}

// Constructor Function
func newUser(firstName string, lastName string, age int, employeeId int, location string) (*user, error) {
	if age == 0 || firstName == "" || lastName == "" || employeeId == 0 || location == "" {
		return nil, errors.New("enter valid inputs")
	}
	return &user{
		firstName,
		lastName,
		age,
		employeeId,
		location,
	}, nil
}

// struct method
func (u *user) printDetails() {

	fmt.Println(u.firstName, u.lastName, u.age, u.employeeId, u.location)
}

func (u *user) modifyDetails(age int, location string) {
	u.age = age
	u.location = location
}

func main() {

	var firstName, lastName, location string
	var age, employeeId int

	fmt.Print("Enter First Name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter EmployeeId: ")
	fmt.Scanln(&employeeId)

	fmt.Print("Enter location")
	fmt.Scanln(&location)

	appUser, err := newUser(firstName, lastName, age, employeeId, location)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	appUser.printDetails()
	appUser.modifyDetails(27, "Bangalore") // mutate state
	appUser.printDetails()
}
