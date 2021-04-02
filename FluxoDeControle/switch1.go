package main

import (
	"fmt"
)

var x interface{}

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

}
