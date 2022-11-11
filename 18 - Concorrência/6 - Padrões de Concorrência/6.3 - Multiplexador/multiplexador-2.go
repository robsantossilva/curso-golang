package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canaisDeEntrada ...<-chan string) <-chan string {
	canalDeSaida := make(chan string)

	for _, canalDeEntrada := range canaisDeEntrada {

		go func(canal <-chan string) {
			for {
				mensagem := <-canal
				canalDeSaida <- mensagem
			}
		}(canalDeEntrada)

	}

	return canalDeSaida

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
