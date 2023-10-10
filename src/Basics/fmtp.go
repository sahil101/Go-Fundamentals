package main

import "fmt"

func main() {
	var name string = "Sahil"

	fmt.Printf("My NAME IS %s\n", name)

	var (
		firstName string
		salary    float64
		gender    bool
	)

	fmt.Printf("My name is %s, I work for cimpress with compensation %.3f, I am a %t", firstName, salary, gender)
}
