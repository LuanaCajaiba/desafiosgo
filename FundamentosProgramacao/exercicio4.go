package main

import "fmt"

var x int = 7

func main() {

	fmt.Printf("%d\n%b\n%#x\n", x, x, x)

	y := x << 1

	fmt.Printf("%d\n%b\n%#x\n", y, y, y)
}
