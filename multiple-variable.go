package main

import "fmt"

func main() {
	//multiple variable
	var firstName, middleName, lastName string
	firstName, middleName, lastName = "sopan", "mukti", "aliyansyah"
	isMariage, age, address := true, 29, "jl raya jagakarsa"
	fmt.Printf("hallo user %s,%s,%s \n", firstName, middleName, lastName)
	fmt.Println("sopan mukti mariage =", isMariage, "age", age, "and", address)
}
