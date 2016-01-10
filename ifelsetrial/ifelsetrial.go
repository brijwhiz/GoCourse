package main

import "fmt"

func main() {
	x, y := 30, 45
	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is equal to y")
	}
}
