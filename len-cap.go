package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	var aaFruits = fruits[0:3]
	fmt.Println(cap(aaFruits))
	fmt.Println(len(aaFruits))
	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
