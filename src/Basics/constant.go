package main

import "fmt"

func main() {
	const days int = 7 // declaring const
	fmt.Println("days", days)

	// days = 8 // cannot assign to days (constant 7 of type int) --> can't assign again to const
	// It is mandatory to assign a value to constant variable
	const x, y = 10, 20 // declaring wihtout type

	const m, n = 10, 0
	// div := m / n -> compile time error because both are constants

	const (
		a = 1090192
		b = 2321
		c
		d
	)

	var e, f = 20, 30
	_, _ = e, f

	// const sum int = e + f // e + f (value of type int) is not constant
	println(a, b, c, d) // 1090192 2321 2321 2321

	const typed float64 = 5
	const untyped = 10

	// const mul int = typed * 5 // throws error

	var mul2 float64 = untyped // implicit conversion
	println(mul2)
	// iota is basically index which is increased by 1 implicitly, number generator for constants
	// x, y = iota, iota  // cannot use iota out of constant declaration

	const val1 = iota
	const val2 = iota

	println(val1, val2) // 0, 0

	const (
		val3 = iota * 2
		val4 = iota * 3
		val5 = iota * 4
	)
	println(val3, val4, val5) // 0 3 8

}
