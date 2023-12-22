package main

import (
	"fmt"
	"math"
)

const ROI float64 = 5.2

func main() {

	var time int
	var principal float64

	fmt.Print("Enter time period: ")
	fmt.Scan(&time)
	fmt.Print("Enter Prinicipal Amt: ")
	fmt.Scan(&principal)

	simpleInterest, compoundInterest := calculateInterest(time, principal)

	fmt.Print("Simple Interest: ", simpleInterest, "\nCompund Interest: ", compoundInterest, " \n")

}

func calculateInterest(time int, principal float64) (float64, float64) {

	simpleInterest := (ROI * float64(time) * principal) / 100.0
	compoundInterest := principal*math.Pow((1+ROI/100.0), float64(time)) - principal

	return simpleInterest, compoundInterest
}
