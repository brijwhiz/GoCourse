package main

import "fmt"

func main() {
	i := 5

	switch i {
	case 4:
		fmt.Println("number is 4")
	case i > 8:
		fmt.Println("i is > 8")
    //this case will cause error due to type mismatch with swithc exp, i
	default:
		fmt.Println("default")
	}
}
