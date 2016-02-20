package main

import "fmt"

func main() {
	var f, c float64

  fmt.Print("Please enter temperature in degrees Fahrenheit: ")
  fmt.Scanf("%f", &f)

	c = (f - 32) * 5 / 9

	fmt.Printf("The equivalent temperature in deg Celsius is: %.2f degrees\n", c)
}
