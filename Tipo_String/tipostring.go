package main

import (
	"fmt"
)

func main() {
	s := "Hello Luana"
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("%v - %T - %#U - %#x", v, v, v, v)
	}

	j := `  colando
	
	acento      grave
	
	tem como          colocar 
	
	bagunçado`
	fmt.Printf("%v \n  %T", j, j)

}
