// programa que calcula a média de uma sala

package main

import "fmt"

func main() {
	lista := []float64{98, 93, 77, 82, 83}
	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	fmt.Println(total / float64(len(lista)))
}
