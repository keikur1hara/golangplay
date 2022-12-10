package main

import (
	"fmt"
)

func main() {
	var t, f bool = true, false
	fmt.Println(t, f)

	var s string = "Hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "test"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println("\"")

	fmt.Println(s[0])
	fmt.Println(string(s[0]))
}
