package main

import "fmt"

func main() {
	var name = make([]string, 2)
	name[0] = "sopan"
	name[1] = "mukti"
	fmt.Println(name)

	var fruits = []string{"banana", "apple", "guava", "papaya"}

	fmt.Println("jumlah element nya \t", len(fruits))
	fmt.Println("isi elemetnya adalah \t", fruits)

	var newFruits = fruits[1:3]
	var aFruits = fruits[0:2]
	var bFruits = fruits[1:3]
	var cFruits = fruits[1:2]
	var dFruits = fruits[0:3]
	fmt.Println(newFruits)
	fmt.Println("afruits", aFruits)
	fmt.Println("bfruits", bFruits)
	fmt.Println("cfruits", cFruits)
	fmt.Println("dfruits", dFruits)
	bFruits[0] = "pinaple"
	fmt.Println("afruits", aFruits)
	fmt.Println("bfruits", bFruits)
	fmt.Println("cfruits", cFruits)
	fmt.Println("dfruits", dFruits)
}
