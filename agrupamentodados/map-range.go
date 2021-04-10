package main

import (
	"fmt"
)

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		198: "não é legal",
		983: "massa",
		846: "Incrível",
	}

	fmt.Println(qualquercoisa)

	//o range não é sempre na ordem

	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}
	fmt.Println(total)

	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)
}
