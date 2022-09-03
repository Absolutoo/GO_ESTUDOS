package main

import "fmt"

func main() {

	var nu1, nu2, result float64
	var op string

	fmt.Println("Calculadora monstra")
	fmt.Println("")
	fmt.Print("Escolha a operação: ")
	fmt.Scan(&op)
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&nu1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&nu2)

	switch op {

	case "+":
		result = nu1 + nu2

	case "-":
		result = nu1 - nu2

	case "/":
		result = nu1 / nu2

	case "*":
		result = nu1 * nu2

	}

	fmt.Println("Resposta: ", result)

}
