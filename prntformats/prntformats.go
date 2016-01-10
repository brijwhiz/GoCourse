package main

import "fmt"

func main() {
	//Formatting booleans
	fmt.Printf("%v\n", false)
	fmt.Printf("%t\n", true)

	//Formatting integers
	fmt.Printf("%v\n", 987)
	fmt.Printf("%d\n", 345)

	//Formatting floats
	fmt.Printf("%v\n", 987123456.987345123432)
	fmt.Printf("%f\n", 987123456.987345123432)
	// Use %g for large exponents
	fmt.Printf("%g\n", 987123456.987345123432)

	/* When formatting numbers you will often wnat to control the width and precision of the resulting figure. To specify the width on an integer, use a number after % in the verb. By default the result will be right-justified and padded with spaces.
	 */
	fmt.Printf("|%6d|%6d|\n", 23, 987)

	/* You can also specify the width of printed floats, though usually you'll also want to restrict the decimal precision at the same with time the width,precision syntax
	 */
	fmt.Printf("|%6.2f|%6.2f|\n", 3.2, 4.37)

	// To left-justify, use the - flag
	fmt.Printf("|%-6.2f|%-6.2f|\n", 3.2, 4.37)

	/* So far we've seen Printf, which prints the formatted string to os.Stdout. Sprintf fromats and returns a string without printing it anywhere
	 */
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Println(2001, 3456)
	fmt.Print(20, 30)
	//no new line at end of Print but new line after Println
}
