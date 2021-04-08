package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5} //número finito especificado
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5} //número não especificado
	fmt.Println(slice)

	slice2 := append(slice, 6) //colocar mais uma posição no slice
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756 //colocando na posição 3 do slice um novo valor
	fmt.Println(slice[3])

}
