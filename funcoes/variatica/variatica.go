package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.f", media(7.5, 8.6, 4.6))
}
