package main

import "fmt"

func main() {
	/*for x := 0; x < 10; x++ { //for x é igual a 0, enquanto x for menor que 10 acrescenta-se 1 ao x
		fmt.Println(x)
	}

	for x := 0; x != 5; x++ {
		fmt.Println(x)
	}

	for x := 0; x <= 5; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println("Impares", x)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for {
		fmt.Println("Loop")
		break
	}*/

	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("x é menor que 10")
		} else {
			fmt.Println("x não é menor que 10")
			break
		}
	}

}

//inicialização x:=0
//condição  x < 10
//pós x++

//na linguagem go ao final de toda instrução existe um ponto e vírgula, se não posto pelo usuário é posto pelo compilador
