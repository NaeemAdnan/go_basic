package main

import "fmt"

func getNumber() int {

	return 10
}

func main() {

	number := getNumber()
	fmt.Println("The number is: ", number)
}