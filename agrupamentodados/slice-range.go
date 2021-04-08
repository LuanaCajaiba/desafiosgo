package main

import "fmt"

func main() {
	slice := []string{"banana", "maça", "jaca"}

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor:", valor)
	}

	slice[3] = "melancia"

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor: ", valor)
	}
}
