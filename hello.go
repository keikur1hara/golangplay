package main

import (
	"fmt"
)

func main() {
	var arri [3]int
	fmt.Println(arri)

	fmt.Printf("%T\n", arri)

	var arri2 [3]string = [3]string{"A", "B"}
	fmt.Println(arri2)

	arri3 := [3]int{1, 2, 3}
	fmt.Println(arri3)

	arri4 := [...]string{"C", "D"}
	fmt.Println(arri4)
	fmt.Printf("%T\n", arri4)

	fmt.Println(arri2[0])
	fmt.Println(arri2[1])
	fmt.Println(arri2[2-1])

	arri2[2] = "C"
	fmt.Println(arri2)

	fmt.Println(len(arri))

}
