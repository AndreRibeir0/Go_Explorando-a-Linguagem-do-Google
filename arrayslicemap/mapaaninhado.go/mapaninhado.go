package main

import "fmt"

func main() {
	funcsProLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsProLetra, "P")

	for letra, funcs := range funcsProLetra {
		fmt.Println(letra, funcs)
	}
}
