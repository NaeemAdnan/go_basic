package main

import "fmt"

 func add(number1 float64, number2 float64) float64 {
 	
	sum := number1 + number2

	return sum
 }

 func main() {
 	a := 5.5
	b := 4.5
	result := add(a, b)
 	fmt.Println("The sum is: ", result)
 }