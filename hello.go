package main

import "fmt"

const Pi = 3.14

const (
	URL      = "http://google.com"
	SiteName = "Google"
)

const (
	A = iota
	B
	C
	D = "A"
	E
	F
)

var Big int

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

}
