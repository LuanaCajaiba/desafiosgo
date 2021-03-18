package main

import (
	"fmt"
)

// foi criado um tipo subjacente a int
type hotdog int

// foi criada uma variável para esse tipo, com o identificador x, utilizando a palavra-chave var
var x hotdog

func main() {
	//foi demonstrada o valor da variável x
	fmt.Println(x)
	//foi demonstrada o tipo de variável x
	fmt.Printf("%T \n", x)
	//foi atribuída 42 a variável x utilizando o operador =
	x = 42
	//foi demonstrado o valor da variável x
	fmt.Println(x)

}
