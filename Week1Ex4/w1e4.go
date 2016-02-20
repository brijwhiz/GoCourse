package main

import "fmt"

func main (){

	type Counter int

	var odometer Counter = 0

	fmt.Printf("The type is %T and the value is %v", odometer, odometer)

}
