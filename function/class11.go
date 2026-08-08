package main

import "fmt"

func getNumbers(num1 int,num2 int) (int, int, float64) {
		sum := num1 + num2
		mul := num1 * num2
		div := float64(num1) / float64(num2)
	return sum, mul, div 
}

func main(){
	a := 5
	b := 10
	sum, mul, div := getNumbers(a, b)
	fmt.Println("The sum is: ", sum)
	fmt.Println("The multiplication is: ", mul)
	fmt.Println("The division is: ", div)
}