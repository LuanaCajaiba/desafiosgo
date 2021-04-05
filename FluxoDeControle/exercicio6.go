package main

import "fmt"

func main() {

	x := 10

	switch {
	case x > 10:
		fmt.Println("Xis é maior que dez ")

	case x == 10:
		fmt.Println("Xis é igual a 10")
	}
}
