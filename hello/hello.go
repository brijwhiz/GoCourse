// single line comment. Line comments start with the
// character sequence // and stop at the end of the line
// A line comment acts like a newline.
// Comments do not nest.

/*
   Multi-
      line comment
*/

/*
Package main.
All Go files start with package <something>.
package main is required for a standalone executable.
*/

package main

import "fmt"

//import "fmt"
// I commented the above code and ran goimports to get the import
// file automatically added to the source code. goimports hello.go

func main() {
	//Mulitple statements separated by ;
	fmt.Printf("Hello ")
	fmt.Printf("World\n")

	// Single statement then ; not needed
	fmt.Printf("Learning Go.\n")
}
