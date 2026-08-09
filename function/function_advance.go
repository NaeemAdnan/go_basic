package main

import "fmt"

func welcome(){
	fmt.Println("Welcome to the app")
}

func getUserInput()string{
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	return name
}

func getNumbers()(int, int){
	var num1 int
	var num2 int

	fmt.Println("Enter your first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter your second number: ")
	fmt.Scanln(&num2)

	return num1, num2
}


func main(){
	welcome()
	name := getUserInput()
	num1, num2 := getNumbers()
	sum := num1 + num2
	fmt.Println("Hello ", name)
	fmt.Println("Sum: ", sum)
}