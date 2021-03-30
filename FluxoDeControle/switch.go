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

}
