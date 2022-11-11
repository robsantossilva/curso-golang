package main

import (
	"fmt"
	"time"
)

func main() {
	receivedChannel := make(chan string)
	defer close(receivedChannel)

	canal := escrever("Olá Mundo!", receivedChannel)

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string, canal chan string) <-chan string {

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
