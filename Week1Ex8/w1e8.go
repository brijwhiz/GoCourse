package main

import "fmt"
import "ftoc"

func main() {

  var f float64
  
	fmt.Print("Please enter temperature in degrees Fahrenheit: ")
	fmt.Scanf("%f", &f)
	fmt.Printf("The equivalent temperature in deg Celsius is: %.2f degrees\n", ftoc.Convertftoc(f))
}
