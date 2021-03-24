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

	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB = 1 << (iota * 10) // 1 << (2 * 10)
		GB = 1 << (iota * 10) // 1 << (3 * 10)
		TB = 1 << (iota * 10) // 1 << (4 * 10)
	)
)

func main() {
	fmt.Println(a, b, c, d)

	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
