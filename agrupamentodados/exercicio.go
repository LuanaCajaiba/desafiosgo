package main

import (
	"fmt"
)

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	for indice, value := range array {
		fmt.Println(indice, value)
	}

	fmt.Printf("%T\n", array)
}
