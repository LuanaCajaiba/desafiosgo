package main

import "fmt"

func main() {

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Luana",
		idade: 23,
	}

	fmt.Println(x)
}
