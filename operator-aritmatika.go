package main

import "fmt"

func main() {
	var val = (((2+6)%3)*4 - 2) / 3

	fmt.Printf("value  adalah  %d\n", val)
	var isEqual = val == 2
	fmt.Println(isEqual)
	fmt.Printf("nilai %d adalah (%t)", val, isEqual)
}
