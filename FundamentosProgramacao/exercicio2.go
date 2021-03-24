package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30
	d := 40
	e := 50
	f := 60

	if c == b {
		fmt.Println("10 é maior que 20")
	} else if c >= d {
		fmt.Println(" 30 é maior ou igual a 40")
	} else {
		fmt.Println("Nem 10 é maior que 20, nem 30 é maior ou igual a 40")
	}

	if a <= e {
		fmt.Println("10 é menor igual a 50")
	} else {
		fmt.Println("10 não é menor igual a 50")
	}

	if e == f {
		fmt.Println("50 é igual de 60")
	} else if e != f {
		fmt.Println("50 é diferente de 60")
	} else {
		fmt.Println("50 não é diferente nem igual a 60")
	}
}
