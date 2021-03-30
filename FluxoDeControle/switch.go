package main

import (
	"fmt"
)

func main() {

	x := 10

	switch {
	case x > 5:
		fmt.Println("Xis é maior que cinco")
	case x == 5:
		fmt.Println("xis é igual a cinco")
	case x < 5:
		fmt.Println("Xis é maior que cinco")

	}

	quemtanoescritorio := "Luana"

	switch quemtanoescritorio {
	case "Luiz":
		fmt.Println("Quem tá no escritório é o Luiz ")
	case "Joana":
		fmt.Println("Quem tá no escritório é a Joana ")
	case "Maria":
		fmt.Println("Quem tá no escritório é a Maria ")
	case "Luana":
		fmt.Println("Quem tá no escritório é a Luana ")

	}

	quemtanoescritorio := "Luan"

	switch quemtanoescritorio {
	case "Luiz", "Carine", "Milena":
		fmt.Println("Quem tá no escritório é o time 1 ")
	case "Joana", "Adão", "Joice":
		fmt.Println("Quem tá no escritório é o time 2 ")
	case "Maria", "Rose", "João":
		fmt.Println("Quem tá no escritório é o time 3 ")
	case "Luana", "Rodrigo", "Andre":
		fmt.Println("Quem tá no escritório é o time 4 ")
	default:
		fmt.Println("Não tem ninguém")

	}

}

//switch avalia o código de cima para baixo.
//se houver uma expressão correta na 1ºexpressão
//ele retornará apenas a 1ºexpressão
