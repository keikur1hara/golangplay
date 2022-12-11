package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I'm a function!")
	})

}
