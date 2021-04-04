package main

import "fmt"

func main() {
	x := 6

	if (x == 2) || (x == 3) {
		fmt.Println("X é dois ou três")
	}

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("é multiplo de dois e também de três")
	}

	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println(" não é multiplo de dois e também de três")
	}
}
