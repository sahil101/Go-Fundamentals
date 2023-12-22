package main

import (
	"fmt"
	"os"
)

func main() {
	// var sentence string = "Hello My Name is Sahil Agarwal, I welcome you to my mock project please enjoy."
	// fmt.Print("Enter a sentence: ")
	// fmt.Scan(&sentence)
	// fileData := fmt.Sprint(sentence)

	// os.WriteFile("mock.txt", []byte(fileData), 0644)

	readData, errors := os.ReadFile("mock.txt")
	// convertedData = strconv
	fmt.Printf(string(readData), errors)
	panic(errors)

}
