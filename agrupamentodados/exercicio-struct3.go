package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	bmw := sedan{veiculo{4, "rosa"}, true}

	carrodeluana := caminhonete{
		veiculo: veiculo{
			portas: 8,
			cor:    "azul",
		},
		traçãoNasQuatro: true,
	}

	fmt.Println(bmw)
	fmt.Println(carrodeluana)

}
