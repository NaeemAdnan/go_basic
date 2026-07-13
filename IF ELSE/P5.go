package main

import "fmt"

func main(){

	a := -5

	if a > 0 {
		fmt.Println("a is positive")
	} else if a < 0{
		fmt.Println("a is negative")
	} else if a == 0{
		fmt.Println("a is zero")
	}
}