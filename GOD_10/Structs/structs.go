package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Absolutoo"
	u.idade = 23
	fmt.Println(u)

	enedrecoExemplo := endereco{"Rua dos Bobos", 0}
	fmt.Println(enedrecoExemplo)

	usuario2 := usuario{"Cleitinho", 30, enedrecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 30}
	fmt.Println(usuario3.idade)

}
