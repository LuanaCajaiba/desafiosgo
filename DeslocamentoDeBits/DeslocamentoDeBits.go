package main

import (
	"fmt"
)

func main() {

	x := 1
	y := x << 1 //operador deslocamento de bits

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
}
