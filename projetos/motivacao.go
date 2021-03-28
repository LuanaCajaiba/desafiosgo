package main

import "fmt"

func main() {

	var x int

	fmt.Println("***************************************************************************")
	fmt.Println("***************************************************************************")
	fmt.Println("***************************************************************************")
	fmt.Println("***************************************************************************")
	fmt.Println("********** Você quer receber a frase motivacional em qual área ************")
	fmt.Println("**********                1 - Financeira                       ************")
	fmt.Println("**********                2 - Acadêmica                        ************")
	fmt.Println("**********                3 - Familiar                         ************")
	fmt.Println("**********                4 - Relacionamento                   ************")
	fmt.Println("***************************************************************************")
	fmt.Println("***************************************************************************")
	fmt.Println("***************************************************************************")
	fmt.Println("***************************************************************************")
	fmt.Scan(&x)

	fmt.Println("Quantas vezes deseja ver")

	if x == 1 {

		fmt.Println("Regra número 1: nunca perca dinheiro. Regra número 2: não esqueça a regra número 1")
	}

	if x == 2 {
		fmt.Println("O aprendiz é um mestra em formação")
	}

	if x == 3 {
		fmt.Println("O que você pode fazer para promover a paz mundial? Vá para casa e ame a sua família")
	}

	if x == 4 {
		fmt.Println("Ame a si mesmo")
	}

}
