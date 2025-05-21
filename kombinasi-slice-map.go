package main

import "fmt"

func main() {

	var students = []map[string]string{
		map[string]string{"name": "sopan", "age": "26"},
		map[string]string{"name": "mukti", "age": "22"},
		map[string]string{"name": "aliyansyah", "age": "21"},
		map[string]string{"name": "agus", "age": "39"},
	}
 
	for _, student := range students {
		fmt.Println("nama siswa", student["name"], "\t umur:", student["age"])
	}
}
