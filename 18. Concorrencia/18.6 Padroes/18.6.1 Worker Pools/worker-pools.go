package main

import (
	"fmt"
	"time"
)

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

// tasks <-chan int canal que somente recebe valores
// results chan<- int canal que somente envia valores
func worker(tasks <-chan int, results chan<- int) {
	for num := range tasks {
		results <- fibonacci(num)
	}
}

func main() {
	momento := time.Now()
	time.Sleep(5 * time.Second)
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	// esse worker quando trabalha sozinho adiciona ao resultado sozinho
	// mas posso chamar outros para ajudar a fazer o trabalho
	// cada um vai adicionar resultados de forma sincrona no canal results
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)


	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
	fmt.Println(".............................................")
	fmt.Println(time.Since(momento))

}
