package main

import "fmt"

func main() {
	fmt.Println("Map")
	var chicken = map[string]int{}
	//chicken = map[string]int{}
	chicken["januari"] = 20
	chicken["februari"] = 22

	fmt.Println(chicken)
	fmt.Println(chicken["januari"])
	fmt.Println(chicken["februari"])
	fmt.Println(chicken["maret"])
	fmt.Println("inisiai nilai map")

	var fruits = map[string]string{
		"banana": "yellow",
		"apple":  "red",
		"grape":  "purple",
		"melon":  "green",
	}

	fmt.Println(fruits["banana"])
	fmt.Println(fruits)

	var fruits2 = map[string]string{}
	var fruits3 = make(map[string]int)
	var fruits4 = *new(map[string]string)
	fmt.Println(fruits2)
	fmt.Println(fruits3)
	fmt.Println(fruits4)

	//iterasi map dengan for range
	for key, value := range fruits {
		fmt.Println(key, "\t", value)
	}
	fmt.Println(len(fruits))
	delete(fruits, "apple")
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var fruit, isExist = fruits["banana"]
	if isExist {
		fmt.Println("fruit banana ada", fruit)
	} else {
		fmt.Println("fruit banana tidak ada")
	}
}
