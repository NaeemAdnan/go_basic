package main

import "fmt"

func main() {
	//Print Welcome to the app
	fmt.Println("Welcome to the app")

	//Get user input for name
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	//Get two number from user

	var num1 int
	var num2 int
	
	fmt.Println("Enter your first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter your second number: ")
	fmt.Scanln(&num2)

	sum := num1 + num2 

	//Display the output

	fmt.Println("Hello ", name)
	fmt.Println("Sum: ", sum)


	//The final thanks and greeting

	fmt.Println("Thank you for using the app")
	fmt.Println("Good Bye")



}