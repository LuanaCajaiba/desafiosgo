package main

import (
	"fmt"
)

func main() {
	/*	nomes := map[string]string{
			"Luana_Cajaiba":   "Programar",
			"Larissa_Cajaiba": "jogar",
		}

		fmt.Println(nomes)

		for i, v := range nomes {
			fmt.Println(i, v)
		}
	*/
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

}
