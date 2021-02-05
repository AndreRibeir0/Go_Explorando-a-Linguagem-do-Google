package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      11325.54,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.02,
	}

	funcESalarios["Rafael Filho"] = 1325.0
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
