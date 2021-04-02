package main

import (
	"fmt"
)

var x interface{}
var y int

func main() {
	x = true
	switch x.(type) {

	case bool:
		fmt.Println(" é um bool ")
	case int:
		fmt.Println(" é um int ")
	case string:
		fmt.Println(" é uma string ")
	case float64:
		fmt.Println(" é um float64 ")

	}

	switch x == 2 {
	case x == 1:
		fmt.Println("é 1")
	case x == 2:
		fmt.Println("é 2")

	}

}
