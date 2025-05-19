package main

import "fmt"

func main() {
	var point = 3
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus nilai anda %d \n", point)
	}

	var value = 4453.0

	if percent := value / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s goods!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

}
