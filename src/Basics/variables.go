package main

import "fmt"

func main() {

	var age int = 30
	var name = "Sahil Agarwal"
	fmt.Println("My age is ", age) // For printing statements to the standard output

	title := "my life is boring" // shorthand property to decalre variables
	fmt.Println("title", title)

	_ = name // Blank identifier act as a black hole, mutes the unused error also can be only on the left hand side of =

	// var salary = 20000
	// salary = 20.4 -> cannot assign int to float, GO is a strongly typed language

	// multiple reassignments using shorthand property only possible if any one variable is new

	// age, name := 40, "ishan" -> this is wrong

	age1, name := 40, "ishan" // multiple assignment
	// := (colon equals syntax) used only when declaring a new variable (or at least a new variable)

	// can't use short declaration at Package Scope (outside main() or other function)
	// all statements at package scope must start with a Go keyword (package, var, import, func etc)

	sum := age + age1 // expressions using short hand

	fmt.Println("Sum is", sum)
	_ = age1

	var (
		firstName string
		salary    float64
		gender    bool
	) // better readability
	fmt.Println(firstName, salary, gender)

	var a, b int // declaration of multiple varaibles of same data type

	fmt.Println(a, b)

}
