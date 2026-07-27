package main

import "fmt"

func getNumbers(num1 int, num2 int) (int, int, int) {

	sum := num1 + num2
	mul := num1 * num2
	div := num1 / num2

	return sum, mul, div

}

func main(){
	a := 20
	b := 10
	p, q, r := getNumbers(a, b)
	fmt.Println("The sum is: ", p)
	fmt.Println("The multiplication is: ", q)
	fmt.Println("The division is: ", r)
}