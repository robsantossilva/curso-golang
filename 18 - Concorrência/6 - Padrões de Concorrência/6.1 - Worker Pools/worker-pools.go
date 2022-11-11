package main

import (
	"fmt"
	"sync"
)

func main() {

	var waitGroup sync.WaitGroup

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	waitGroup.Add(4)

	go worker(tarefas, resultados, &waitGroup)
	go worker(tarefas, resultados, &waitGroup)
	go worker(tarefas, resultados, &waitGroup)
	go worker(tarefas, resultados, &waitGroup)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	// for i := 0; i < 45; i++ {
	// 	resultado := <-resultados
	// 	fmt.Println(resultado)
	// }

	go func() {
		for resultado := range resultados {
			fmt.Println(resultado)
		}
	}()

	waitGroup.Wait()
	close(resultados)

}

func worker(tarefas <-chan int, resultados chan<- int, waitGroup *sync.WaitGroup) {
	//waitGroup.Add(1)
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
	waitGroup.Done()
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
