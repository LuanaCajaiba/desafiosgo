package main

import (
	"fmt"
)

//slice de slice

func main() {
	ss := [][]string{
		[]string{
			"Luana",
			"Cajaiba",
			"Programar",
		},
		[]string{
			"Luciano",
			"Hulk",
			"Andar de lancha",
		},
		[]string{
			"Maria",
			"Gadú",
			"Cantar",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Println("\n\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}

}
