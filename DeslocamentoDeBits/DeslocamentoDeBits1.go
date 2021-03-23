package main

import (
	"fmt"
)

func main() {

	y := 24

	x := y << 2 // deslocamento em duas casas
	z := y >> 2

	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)

}
