//Escreva um código em Go utilizando todo o conhecimento
//adquirido até o momento. E nesse código você precisará, baseado
//em nossas aulas de concorrência, utilizar canais e goroutines
//para que o seu programa exiba, em alternância,
//as palavras ping e pong.

package main

import (
	"fmt"
	"time" //para imprimir uma msg várias vezes com um intervalo de tempo entre elas
)

func pingar(y chan string) { // palavra reservada para canal: chan
	for i := 0; ; i++ {
		y <- "ping" //<- símbolo usado para enviar e receber mensagem pelo canal
	}
}

func pongar(z chan string) {
	for i := 2; ; i++ {
		z <- "pong"
	}
}

func imprimir(y chan string) {
	for {
		msg := <-y
		fmt.Println(msg)
		time.Sleep(time.Second * 1) //intervalo de tempo
	}
}

func main() {
	var y chan string = make(chan string)
	go pingar(y)
	go pongar(y)
	go imprimir(y)
	var input string
	fmt.Scanln(&input)
}
