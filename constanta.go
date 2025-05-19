package main

import "fmt"

func main() {
	const fistName string = "Sopan"
	const lastName string = "Mukti"

	fmt.Print("hallo firstname"+" ", fistName+"\n")
	fmt.Print("hallo lastname"+" ", lastName+"\n")

	//multiple constanta
	fmt.Println("-------multple constanta-------------")
	const (
		square        = "kotak"
		isToday  bool = true
		numeric  uint = 2
		floatNum      = 2.2
	)
	fmt.Println(square, isToday, numeric, floatNum)
	//jika variable di deklarasikan tanpa memberikan type data dan nilai maka ia ikut sama dengan variable di atasnya
	const (
		a = "konstanta"
		b
	)
	fmt.Printf("%s %s", a, b)
	const (
		today = "sekarang"
		sekarang
		isToday2 = true
	)
	fmt.Println(today, sekarang, isToday2)
	const one, two = 1, 2

	const job, work string = "programer", "programing"

	fmt.Println(one, two, job, work)
}
