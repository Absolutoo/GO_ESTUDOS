package main

import "fmt"

func main() {
	//ARITIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	mult := 10 * 5
	restoD := 10 % 2

	fmt.Println(soma, subtracao, divisao, mult, restoD)

	var n1 int = 10
	var n2 int = 25
	soma2 := n1 + n2
	fmt.Println(soma2)

	// FIM DOS ARITIMETICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// FIM DOS OPERADORES DE ATRIBUIÇAO

	// OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// FIM DOS RELACIONADOS

	// OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numer0 = numeor -20

	numero *= 3 //numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
