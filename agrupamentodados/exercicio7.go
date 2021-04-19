package main

import (
	"fmt"
)

func main() {
	mepe := map[string][]string{
		"Souza_Orlando": []string{
			"dormir", "comer", "trabalhar",
		},
		"Junqueira_Tereza": []string{
			"programar", "andar-de-moto", "surfar",
		},
		"Huck_Luciano": []string{
			"apresentar-programa", "dar-casa-reformada", "se-candidatar-presidente",
		},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
	}

	mepe["Albuquerque_Rogério"] = []string{
		"jogar-futebol", "jogar-mario", "assistir",
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
	}

}
