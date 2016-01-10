package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x) // Prints 1

	{
		fmt.Println(x) // Prints 1
		x := 2         // looks like assignment statement in other lang, but is declaration
		fmt.Println(x) // Prints 2
	}

	fmt.Println(x) // Prints 1 (bad if you need 2)
}
