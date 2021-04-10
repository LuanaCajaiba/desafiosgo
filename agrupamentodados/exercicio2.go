package main

import (
	"fmt"
)

func main() {
	slice := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	numeros0 := slice[0:3]
	fmt.Println(numeros0)

	numeros1 := slice[4:]
	fmt.Println(numeros1)

	numeros2 := slice[1:7]
	fmt.Println(numeros2)

	numero3 := slice[2:9]
	fmt.Println(numero3)
}
