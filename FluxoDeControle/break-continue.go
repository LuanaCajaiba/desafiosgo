package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			//continua quando o núemro não é par
			continue //quando o número não for par eu paro a execução dessa iteração e vou direto para a próxima
		}
		fmt.Println(i)
	}

	for i := 0; i <= 20; i++ {
		if i == 3 {
			continue
			//aqui não será mostrado o 3, pulará para o próximo
		}
		fmt.Println(i)
	}

}

//o break pega o loop inteiro e joga fora
//o continue quebra apenas uma iteração específica
