package main

import (
	"fmt"
)

const (
	a = iota * 10
	b
	_ //usado para pular o número
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)

}
