package main

import (
	"fmt"
)

func main() {
	s := "Hello Luana"
	sb := []byte(s)

	fmt.Printf("%v \n %T", sb, sb)

	j := `  colando
	
	acento      grave
	
	tem como          colocar 
	
	bagunçado`
	fmt.Printf("%v \n  %T", j, j)

}
