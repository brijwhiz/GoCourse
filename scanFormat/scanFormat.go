package main

import "fmt"

func main() {
	var username string

	fmt.Println("Enter username: ")
	fmt.Scanf("%s", &username)

	fmt.Println("Username:", username)
}
