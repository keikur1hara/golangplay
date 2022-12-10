package main

import (
	"fmt"
)

func outer() {
	var s4 = "outer"
	fmt.Println(s4)
}
func main() {
	fmt.Println("Hello, World")

	var i int = 100

	fmt.Println(i)

	var s string = "Hello go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	i4 := 400
	fmt.Println(i4)

	outer()

}
