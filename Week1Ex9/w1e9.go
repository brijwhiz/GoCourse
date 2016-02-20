package main

import "fmt"

func main() {
	switch {
	default:
		fmt.Println("I am default")
	case false:
		fmt.Println("I am false")
	case true:
		fmt.Println("I am true")
	}
}
