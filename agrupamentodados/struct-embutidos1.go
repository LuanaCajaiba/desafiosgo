package main

import (
	"fmt"
)

type pessoa struct {
	pessoa
	título  string
	salario int
}

func main() {

	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.salario)

}
