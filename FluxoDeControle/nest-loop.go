package main

import "fmt"

//repetição hierárquica
func main() {
	for hora := 0; hora <= 12; hora++ {
		fmt.Println("h: ", hora)

		for minuto := 0; minuto < 60; minuto++ { //esse loop será repetido inteiro a cada for rodado em cima
			fmt.Print(" ", minuto)
		}
		fmt.Println("\n")
	}
}
