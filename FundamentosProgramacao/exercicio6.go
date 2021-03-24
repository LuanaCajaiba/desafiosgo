package main

import "fmt"

const (
	a = iota + 2022
	b
	c
	d
)

const (
	_ = 2050 + iota
	e
	f
	g
	h
)

func main() {

	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)

}
