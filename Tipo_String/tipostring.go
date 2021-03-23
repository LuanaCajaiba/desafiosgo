package main

import (
	"fmt"
)

func main() {
	s := "Hello Luana"
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
		//v - valor numérico
		//t - tipo
		//unicode
		//tipo hexadecimal
	}

	fmt.Println(" ")
	//o resultado será caractere por caractere - UTF8
	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}
	//o resiltado será byte por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}

	fmt.Println(" ")

	j := `  colando
	
	acento      grave
	
	tem como          colocar 
	
	bagunçado`
	fmt.Printf("%v \n  %T", j, j)

}
