package main

import "fmt"

func main() {

	age := 20

	if age > 20 {
		fmt.Println("You are eligible for marriage")
	} else  if age == 20{
		fmt.Println("You are not eligible for marriage, but you can start looking for a partner")
	} else {
		fmt.Println("You are not eligible for marriage")
	}
}
