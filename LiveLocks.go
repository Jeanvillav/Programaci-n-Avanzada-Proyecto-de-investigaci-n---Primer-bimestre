package main

import (
	"fmt"//funciones de entrada/salida como Println
	"math/rand"
	"sync"//Proporciona primitivas de sincronización como WaitGroup
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	wg.Add(2)

	resource1 := make(chan struct{}, 1)
	resource1 <- struct{}{}
	resource2 := make(chan struct{}, 1)
	resource2 <- struct{}{}
	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {//La goroutine intentará adquirir ambos recursos hasta 5 veces.
			for {
				select {
				case <-resource1:// Intentar adquirir resource1
					select {
					case <-resource2:// Intentar adquirir resource2
						fmt.Println("Goroutine 1 adquirió ambos recursos")
						resource1 <- struct{}{} // Liberar resource1
						resource2 <- struct{}{} // Liberar resource2
						return
					default:
						fmt.Println("Goroutine 1 liberando resource1")
						resource1 <- struct{}{}//Si resource2 no está disponible, libera resource1
					}
				default:
					fmt.Println("Goroutine 1 esperando")
					time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)// Esperar antes de reintentar
				}
			}
		}
	}()//Cierra la declaración de la goroutine

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			for {
				select {
				case <-resource2:// Intentar adquirir resource2
					select {
					case <-resource1: // Intentar adquirir resource1
						fmt.Println("Goroutine 2 adquirió ambos recursos")
						resource2 <- struct{}{}// Liberar resource2
						resource1 <- struct{}{}// Liberar resource1
						return
					default:
						fmt.Println("Goroutine 2 liberando resource2")
						resource2 <- struct{}{} // Liberar resource2
					}
				default:
					fmt.Println("Goroutine 2 esperando")
					time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)// Esperar antes de reintentar
				}
			}
		}
	}()

	wg.Wait() // Esperar a que las goroutines terminen
}
