package main

//key is the . alias for the package.
//that allows us to remove the explicit package refernce

import . "fmt"

// the go-plus package of atom is shouting about using the dot imports

func main() {
	Println("hello world")
}
