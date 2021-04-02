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

	quemtanoescritorio1 := "Luan"

	switch quemtanoescritorio1 {
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

	a := 1

	switch a {
	case 1, 2:
		fmt.Println("1 ou 2 ")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("Não tem ninguém")

	}

	z := 4

	switch {
	case (z == 4), (z == 8):
		fmt.Println("Z é igual a 4 ou a 8")
	case (z < 10), (z < 11):
		fmt.Println("Z é menor que 10 ou menor que 11")
	}

}

//switch avalia o código de cima para baixo.
//se houver uma expressão correta na 1ºexpressão
//ele retornará apenas a 1ºexpressão
