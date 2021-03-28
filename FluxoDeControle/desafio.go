package main

import (
	"fmt"
)

func main() {

	for i := 33; i <= 122; i++ {

		fmt.Printf("%#U \t ", i)
	}

	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v\n", x, string(x))
	}
}
