package main

import "fmt"

func printsomething(){
	fmt.Println("Hello World")
}


func sayHello(name string){
	fmt.Println("Hello ", name)
}

func main(){
	printsomething()
	sayHello("Naeem")
}