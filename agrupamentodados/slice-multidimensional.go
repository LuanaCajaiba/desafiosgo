package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		//indice      0  1  2   Indice
		[]int{1, 2, 3}, //0
		[]int{4, 5, 6}, //1
		[]int{7, 8, 9}, //2
	}

	fmt.Println(ss)
	fmt.Println(ss[1])
	fmt.Println(ss[1][1])
}
