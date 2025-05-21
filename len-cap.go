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

	var cFruits = fruits[0:2]
	fmt.Println(len(cFruits))
	fmt.Println(cap(cFruits))
	var dFruits = append(cFruits, "manggo")
	fmt.Println(dFruits)
	fmt.Println(cFruits)
	fmt.Println(fruits)

	//fungsi copy
	fmt.Println("fungi copy")
	var src = make([]string, 5)
	var dst = []string{"apple", "grape", "banana", "melon", "mango", "orange"}
	var n = copy(src, dst)

	fmt.Println("N", n)
	fmt.Println("dst", dst)
	fmt.Println("src", src)

	var src2 = []string{"apple", "apple", "apple", "apple"}
	var dst2 = []string{"banana", "grape", "melon"}
	var v = copy(src2, dst2)

	fmt.Println("dst2", dst2)
	fmt.Println("src2", src2)
	fmt.Println("v", v)

	//pengaksesan 3 index pada slice

	fmt.Println("pengaksesan 3 index pada slice")
	var fruitss = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruitss)
	fmt.Println(len(fruitss))
	fmt.Println(cap(fruitss))
	var a = fruits[0:3]
	var b = fruits[0:3:3]
	fmt.Println(a)
	fmt.Println("a", len(a))
	fmt.Println("a", cap(a))
	fmt.Println(b)
	fmt.Println("b", len(b))
	fmt.Println("b", cap(b))

}
