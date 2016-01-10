package main

import "fmt"

// myvar := 1 short variable declartion can be used only within functions.

func main() {
	i := 0
	//i := 2
	// error : no new variables on left side of :=

	i, j := 1, 2
	//the redeclaration of i is possible because it is not alone in the
	// second line

	i, j = j, i
	fmt.Println(i, j)
}
