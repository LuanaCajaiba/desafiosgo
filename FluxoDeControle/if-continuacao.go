package main

import "fmt"

func main() {
	x := 10
	if x < 100 {
		fmt.Println("xis é menor que cem")
	} else {
		fmt.Println("xis não é menor que cem")
	}

	a := 10
	if a < 100 {
		fmt.Println("a é menor que cem")
	} else if a > 10 {
		fmt.Println("a é maior que dez")
	}

	if a < 100 {
		fmt.Println("a é menor que cem")
	} else if a > 10 {
		fmt.Println("a é maior que dez")
	} else {
		fmt.Println("a não é nem menor, nem maior que 10")
	}
}
