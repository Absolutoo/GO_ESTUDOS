package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 100
	fmt.Println(numero)
	var numero1 uint = 1000000
	fmt.Println(numero1)

	// alias
	//uint32 não deixa colocar numeros negativos
	var numero3 rune = 325
	fmt.Println(numero3)

	//uint8
	var numero4 byte = 3
	fmt.Println(numero4)

	var numero5 float32 = 123.4
	fmt.Println(numero5)

	var numero6 float64 = 1321.13412
	fmt.Println(numero6)

	//FIM NUMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Testo2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//FIM STRINGS

	var buling bool
	fmt.Println(buling)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
