//Criar uma Calculadora completa:
//a função de soma, subtração, adição e multiplicação e apresente os resultados corretos na tela.

package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	var ope string
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scanln(&x)
	fmt.Print("Digite a operação que deseja (ex: + - ** / *): ")
	fmt.Scanln(&ope)
	fmt.Print("Digite o segundo valor: ")
	fmt.Scanln(&y)
	var sum int = maths(x, ope, y)
	fmt.Println(x, ope, y, "=", sum)

}
func maths(x int, ope string, y int) int {
	if ope == "+" {
		return x + y
	} else if ope == "-" {
		return x / y
	} else if ope == "x" || ope == "*" {
		return x * y
	} else if ope == "/" {
		return x / y
	} else if ope == "^" || ope == "**" {
		return x ^ y
	}
	return 1
}
