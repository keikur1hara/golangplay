package main

import (
	"fmt"
)

func main() {
	var i int = 100
	var i2 int64 = 200

	fmt.Println(i + 50)

	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
	fmt.Print(int32(i) + int32(i2))

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Printf("%T\n", fl)
	fmt.Printf("%T\n", fl64)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	zero := 0.0
	pinf := 1.0
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
