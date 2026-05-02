package main

import "fmt"

func add(a int, b int) {
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)
}

func main() {
	a := 5
	b := 10
	add(a, b)

}
