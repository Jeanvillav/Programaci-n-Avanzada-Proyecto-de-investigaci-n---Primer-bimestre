package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup // Se declara un waitgroup
)

func increment() {
	defer wg.Done() // Asegura que wg.Done() se llame al final de la función

	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	wg.Add(2) // Indica que estamos esperando a que dos goroutines terminen

	go increment()
	go increment()

	wg.Wait() // Espera a que ambas goroutines terminen

	fmt.Println("Contador:", counter)
}
