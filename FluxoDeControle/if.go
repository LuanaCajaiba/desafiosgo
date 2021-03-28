package main

import "fmt"

func main() {
	x := 10
	if x < 100 {
		fmt.Println("x é menor que 100")
	}

	if x < 100 {
		fmt.Println("x é menor que 100, só que a condição está entre parenteses")
	}

	if !(x < 100) {
		fmt.Println("x não é menor que 100")
	}

	if !(x > 100) {
		fmt.Println("x não é maior que 100")
	}

}
