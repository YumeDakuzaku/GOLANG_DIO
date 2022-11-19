// Um professor do ensino médio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da água de Kelvin para graus Celsius.
//Dica: C = K-273

package main

import "fmt"

const tempkelvin = 320.0

func main() {
	var K float64 = tempkelvin
	var C float64 = (K - 273) //conversão de Kelvin para Celsius

	fmt.Println("A temperatura convertida de Kelvin para graus Celsius é:", C)
}
