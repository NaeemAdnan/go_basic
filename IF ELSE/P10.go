package main
import "fmt"

func main() {
	Marks := 40
	
	if Marks >= 80 {
	    fmt.Println("A+")
	} else if Marks >= 70{
	    fmt.Println("A")
	} else if Marks >= 60 {
	    fmt.Println("A-")
	} else {
	    fmt.Println("Fail")
	}
}