package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10) //5 elementos e 10 de capacidade

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice)) //a slice, comprimento da slice, capacidade da slice
	//jogou a array fora, criou um novo array com mais elementos, copiou todos os elementos para lá
}
