package main

import "fmt"

func main() {

	const secretKey = 23232390 // Good
	const SecretKey = 23232309 // Good

	//	const secret_key = 23232390 // Bad
	const SECRETKEY = 23232390 //Bad

	// Untyped constants
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point
	const hi = "Hello"  // kind: string

	//Typed constants
	const ti int = 12345 //kind: Integer with int64 precision restrictions
	const tf float64 = 3.14152

	fmt.Println(secretKey, SecretKey /*secret_key,*/, SECRETKEY, ui, uf, hi, ti, tf)
}
