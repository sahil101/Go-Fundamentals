package main

import "fmt"

func main() {

	// loops syntax
	for i := 0; i < 20; i++ {

		// if-else control statement
		if i%2 == 0 {
			fmt.Println("Even Number: ", i)
		} else {
			fmt.Println("Odd Number: ", i)
		}
	}

	// for { ....statements   } - this will create a infinite for-loop
	// for condition { ....statements   } - this will execute until the condition remains true

	// end the function using a empty return statement
	return

	// break, continue works as same

	fmt.Println("Hello! we will never get printed")

	// we also have "switch" statements in which we don't break to stop executing
}
