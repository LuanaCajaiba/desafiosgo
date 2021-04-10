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

	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}
}
