package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"Nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["Nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
