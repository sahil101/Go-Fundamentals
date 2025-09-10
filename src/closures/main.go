package main

import "fmt"

func createClosure(closureName string) func() {
	var counter int32 = 0
	var name string = closureName
	return func() {
		counter++
		fmt.Println(name, counter)
	}
}

func main() {
	closure1 := createClosure("closure1")
	closure2 := createClosure("closure2")
	closure3 := createClosure("closure3")
	closure1()
	closure1()
	closure2()
	closure2()
	closure1()
	closure3()
}
