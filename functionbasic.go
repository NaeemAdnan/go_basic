package main


import "fmt"

func add(num1 int, num2 int){
	
	sum := num1 + num2

	fmt.Println("The sum of two numbers is: ", sum)
	
	//fmt.Println("The sum of two numbers is: ", num1 + num2)
}


func main(){

	a := 20
	b := 10

	add(a, b)

	add(3, 4)
}
