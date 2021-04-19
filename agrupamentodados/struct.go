package main

import (
	"fmt"
)

//struct é um tipo de dado, portanto deve ser declarado
type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {

	cliente1 := cliente{
		nome:      "José",
		sobrenome: "Cajaiba",
		fumante:   false,
	}

	cliente2 := cliente{"maria", "Ritch", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
