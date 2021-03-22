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

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

	fmt.Println(" ")

	j := `  colando
	
	acento      grave
	
	tem como          colocar 
	
	bagunçado`
	fmt.Printf("%v \n  %T", j, j)

}
