package main

import "fmt"

func main() {
	//                      0.         1.          2.      3.            4.
	sabores := []string{"Queijo", "Calabresa", "Atum", "Portuguesa", "Frango"}

	//fatia := sabores[0:2] //picote na cabeça do 2, não está incluido

	//	fatia := sabores[2:len(sabores)]  //aqui pegará do item 2 até o final do slice
	//  fatia := sabore[2: ] //dessa maneira irá do 2 até o final
	//  fatia := sabores[:2] //dessa maneira irá do 0 até o 2

	//fatia := sabores[:]

	//	fmt.Println(fatia)

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}
