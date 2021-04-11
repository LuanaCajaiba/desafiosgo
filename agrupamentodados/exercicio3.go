package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for indice, valor := range x {
		fmt.Println(indice, valor)
	}

	//x[10] = 52 - está errado

	x = append(x, 52)
	for indice, valor := range x {
		fmt.Println(indice, valor)
	}

	x = append(x, 53, 54, 55)
	for indice, valor := range x {
		fmt.Println(indice, valor)
	}

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

}
