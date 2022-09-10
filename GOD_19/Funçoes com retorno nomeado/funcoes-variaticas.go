package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}

func main() {
	fmt.Println("Funções variaticas")
	totalDaSoma := soma(1, 2, 3, 0)
	fmt.Println(totalDaSoma)

	soma2 := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	escrever("Olá mundo", soma2) //Acada número que eu colocar ele cria uma linha

}
