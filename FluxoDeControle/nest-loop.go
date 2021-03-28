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

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês", mes)
		for dias := 1; dias <= 30; dias++ {
			fmt.Print("Dia: ", dias, ", ")
		}
	}
	fmt.Println("\n")
}
