package main

import "fmt"

func main() {
	fmt.Println("Funções anonimas")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)

	}("Passando Parâmetro")
	fmt.Println(retorno)
}
