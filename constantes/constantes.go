package main

import (
	"fmt"
)

const x = 10 //ficará indefinido até ser usado (definido no momento do uso)
//var y = 10 //no momento que o programa rodar, será decidido que é um int (definido no momento da atribuição)
//var c int

var d float64

func main() {

	d = x
	fmt.Println(d)

}
