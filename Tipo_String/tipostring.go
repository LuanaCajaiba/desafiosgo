package main

import (
	"fmt"
)

func main() {
	s := "Hello Luana"
	fmt.Printf("%v \n %T", s, s)

	j := `  colando
	
	aspas             simples
	
	tem como          colocar 
	
	bagunçado`
	fmt.Printf("%v \n  %T", j, j)

}
