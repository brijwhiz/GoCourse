package main

import "fmt"

func main() {
	loop := 1

L:
	fmt.Println(loop)
	loop++
	if loop <= 5 {
		goto L
	}
}
