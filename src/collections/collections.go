package main

import "fmt"

func main() {
	var prices = [4]float64{20.0}
	prices = [4]float64{20.09, 10}
	fmt.Println(prices)

	var age = [6]int{10, 40, 30, 99, 1023, 21321}

	slicedAge := age[:1]
	slicedageAgain := age[1:]
	slicedAga3 := slicedageAgain[0:]

	fmt.Println(cap(slicedAge), len(slicedageAgain), slicedAga3)

	dynamicArrays := []float64{10.99, 10}

	// dynamicArrays[2] = 10 // out of inde	x range error
	// memory management is done by go
	newSlicedArray := append(dynamicArrays, 2900.0)
	fmt.Println(newSlicedArray)
	fmt.Println(dynamicArrays)
	exercise()
}

func exercise() {
	var newArray = []string{"cricket", "books", "coding"}
	fmt.Println((newArray))
	fmt.Println(newArray[0], newArray[1:3])

	newSlice := newArray[0:2]
	fmt.Println(newSlice)
	newSlice = newSlice[1:3]

	fmt.Println(newSlice)
	dynamicSlice := []string{"fundamentals", "interfaces"}

	dynamicSlice[1] = "generices"

	newDynamicSlice := append(dynamicSlice, "pointers")
	fmt.Println(newDynamicSlice)

	// combine two slices
	combinedSlice := append(newArray, newDynamicSlice...)
	fmt.Println(combinedSlice)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
