package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5} //número finito especificado
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5} //número não especificado
	fmt.Println(slice)

	slice2 := append(slice, 6)

	fmt.Println(slice2)

}
