package User

import (
	"errors"
	"fmt"
)

type User struct {
	firstName  string
	lastName   string
	age        int
	employeeId int
	location   string
}

// struct embedding
type Admin struct {
	email    string
	password string
	User
}

// Constructor Function
func New(firstName string, lastName string, age int, employeeId int, location string) (*User, error) {
	if age == 0 || firstName == "" || lastName == "" || employeeId == 0 || location == "" {
		return nil, errors.New("enter valid inputs")
	}
	return &User{
		firstName,
		lastName,
		age,
		employeeId,
		location,
	}, nil
}

func NewAdmin(email string, password string, user User) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User:     user,
	}
}

// struct method
func (admin *Admin) PrintDetails() {

	fmt.Println(admin.firstName, admin.lastName, admin.age, admin.employeeId, admin.location, admin.email, admin.password)
}

func (u *User) ModifyDetails(age int, location string) {
	u.age = age
	u.location = location
}
