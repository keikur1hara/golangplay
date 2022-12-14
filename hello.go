package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println(("Subloop"))
	time.Sleep(100 * time.Millisecond)
}
func main() {

	fmt.Println("Test")
}
