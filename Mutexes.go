package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex // Mutex para proteger el acceso a counter
)

func increment() {
	mu.Lock() // Bloquea el mutex para asegurar exclusión mutua
	counter++
	mu.Unlock() // Desbloquea el mutex para permitir el acceso a otras goroutines
}

func main() {
	go increment()
	go increment()
	
	fmt.Println("Counter:", counter)
}
