package main

import "fmt"

func main(){

	a := 15

	if a > 15 {
		fmt.Println("a is greater than 15")
	} else if a < 15{
		fmt.Println("a is less than 15")
	} else if a == 15{
		fmt.Println("a is equal to 15")
	}
}