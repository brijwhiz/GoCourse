package main

import "fmt"

func main() {

	for num:= 1; num <= 100; num ++ {
		if num%2 != 0 {
			continue
		}

		fmt.Println(num)
	}
}
