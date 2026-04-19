package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutines

	// duas ou mais acessoando o mesmo dado ao mesmo tempo
	// pelo menos um deles está escrevendo
	
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	m2 := make(map[int]int)

	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i
		}
	}()
	go func() {
		for i := 1001; i < 2000; i++ {
			fmt.Println(m2[i])
		}
	}()
	time.Sleep(5 * time.Second)

}
