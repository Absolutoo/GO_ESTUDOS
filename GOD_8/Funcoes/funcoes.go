package main

import "fmt"

func soma(n1 int, n2 int) int {
	return n1 + n2

}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	soma := soma(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	// _ esse tracinho serve para usar uma variavel que eu não quero usar mas que sou obrigado a usar
	// a ordem dos fatores altera o resultado
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
