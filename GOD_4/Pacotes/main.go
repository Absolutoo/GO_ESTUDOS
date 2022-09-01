package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da Main")
	auxiliar.Escrevendo()

	erro := checkmail.ValidateFormat("vinii.alves@hotmail.com")
	fmt.Println(erro)

}
