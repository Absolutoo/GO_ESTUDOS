package main

import "fmt"

func diaDaSenama(numero int) string {
	switch numero { // switch {
	case 1: // case 1
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número Invalido"
	}

}

func diaDaSemana2(numero int) string {
	var dia2 string
	switch {
	case numero == 1:
		dia2 = "Domingo"
		fallthrough
	case numero == 2:
		dia2 = "Segunda-Feira"
	case numero == 3:
		dia2 = "Terça-Feira"
	case numero == 4:
		dia2 = "Quarta-Feira"
	case numero == 5:
		dia2 = "Quinta-Feira"
	case numero == 6:
		dia2 = "Sexta-Feira"
	case numero == 7:
		dia2 = "Sabádo"
	default:
		dia2 = "Número Inválido"
	}

	return dia2
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSenama(200)
	fmt.Println(dia)

	fmt.Println("--------")

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

}
