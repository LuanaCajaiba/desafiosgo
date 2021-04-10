package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"julia":  8758475,
		"thiago": 9746493,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["julia"])

	//colocando números novos

	amigos["gopher"] = 77777

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	//o ,ok determina se existe ou não. o 0 não é o valor de fantasma e só para determinar que não existe valor
	//sera, ok := amigos["fantasma"]
	//fmt.Println(sera, ok)

	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(sera)
	}

}
