package main

import "fmt"

func main() {
	//pengulanan for pada go
	for i := 5; i >= 0; i-- {
		fmt.Println("perulangan ke ", i)
	}
	//perulangan for dengan konsep while
	var j = 0
	for j < 5 {
		fmt.Println("perulangan J ke ", j)
		j++
	}

	//perulangan for dengan konsep tampa argumen

	var s = 0

	for {
		fmt.Println("perulangan s ke ", s)
		s++
		if s == 5 {
			break
		}
	}
	//for range
	fmt.Println("for range")
	var xs = "123"

	for e, v := range xs {
		fmt.Println("index", e, "valuenya", v)
	}

	var ys = [5]int{10, 20, 23, 40, 50}

	for _, s := range ys {
		fmt.Println("value=", s)
	}

	fmt.Println("for range dari slice")
	var zs = ys[0:2]
	for _, p := range zs {
		fmt.Println("value", p)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2}
	for k, v := range kvs {
		fmt.Println("index", k, "valuenya", v)
	}

	fmt.Println("perulangan bersarang")
	for m := 0; m < 5; m++ {
		for s := m; s < 5; s++ {
			fmt.Print(s, " ")
		}
		fmt.Println()
	}

}
