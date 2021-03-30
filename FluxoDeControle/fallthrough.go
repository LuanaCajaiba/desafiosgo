package main

import (
	"fmt"
)

func main() {
	quemtanoescritorio := "ningém"

	switch quemtanoescritorio {

	case "Luiz":
		fmt.Println("Quem está é o Luiz")
		fallthrough
	case "Luana":
		fmt.Println("Sempre que luiz está Luana está")
	case "maria":
		fmt.Println("Quem está é a maria")
	case "Joana":
		fmt.Println("Quem está a Joana")
	default: //se nenhum caso for suprido
		fmt.Println("Tá vazio.")

	}

}

//fallthrough pula a comparação posterior e vai direto para o statement fmt.Println
