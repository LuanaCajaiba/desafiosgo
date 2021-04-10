package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"julia":  8758475,
		"thiago": 9746493,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["julia"])
}
