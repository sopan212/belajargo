package main

import "fmt"

func main() {
	var numbers = [2][3]int{[3]int{3, 2, 5}, [3]int{1, 2, 4}}
	var numbers2 = [2][3]int{{3, 3, 1}, {8, 6, 0}}
	fmt.Println(numbers2)
	fmt.Println(numbers)

	fmt.Println("perulangan array menggunakan for")
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index of %d : %s\n", i, fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("index dari %d = %s\n", i, fruit)
	}
}
