package main

import "fmt"

func main() {
  type MyString string
  var m MyString
  const hi = "Hi"
  const hello string = "Hello"
  const ole MyString = "Ole"


  fmt.Println("hi:", hi)
  fmt.Println("hello:", hello)
  fmt.Println("Ole:", ole)

  m = hi
  fmt.Println("m :", m)

  // Error m = hello
  fmt.Println("m :", m)

  m = ole
  fmt.Println("m :", m)
}
