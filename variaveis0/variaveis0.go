package main

import (
	"fmt"
)

// foi criado um tipo subjacente a int
type hotdog int

// foi criada uma variável para esse tipo, com o identificador x, utilizando a palavra-chave var
var x hotdog

var y int

func main() {
	//foi demonstrada o valor da variável x
	fmt.Println(x)
	//foi demonstrada o tipo de variável x
	fmt.Printf("%T \n", x)
	//foi atribuída 42 a variável x utilizando o operador =
	x = 42
	//foi demonstrado o valor da variável x
	fmt.Println(x)
	//foi utilizada a conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente
	y = int(x)
	// foi demonstrado o valor de y
	fmt.Println(y)
	// foi demonstrado o tipo de y
	fmt.Printf("%T\n", y)

}
