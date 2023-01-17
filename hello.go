package main

import "fmt"

func main() {

	m := map[string]int{
		"Apple":  100,
		"banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(m[k])
	}

}
