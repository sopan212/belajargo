package main

import "fmt"

func main() {
	var right = true
	var left = false

	var leftAndRight = left && right
	fmt.Printf("left && right \t (%t) \n", leftAndRight)
	var leftOrRight = left || right
	fmt.Printf("left || right \t (%t) \n", leftOrRight)
	var leftReverse = !left
	fmt.Printf("!left \t \t (%t) \n", leftReverse)
}
