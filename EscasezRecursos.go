package main

import (
	"fmt"
	"sync" // Paquete que proporciona primitivas de sincronización como WaitGroup
)

func main() {
	const maxWorkers = 10 // Definimos el número máximo de workers permitidos

	jobs := make(chan int, 100) // Creamos un canal de trabajos con un buffer de tamaño 100

	wg := sync.WaitGroup{} // Creamos una instancia de WaitGroup

	// Iniciamos un bucle para crear maxWorkers goroutines (hilos de trabajo)
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1) // Incrementamos el contador interno de WaitGroup en 1
		// Creamos una nueva goroutine
		go func() {
			defer wg.Done() // Se llamará a wg.Done() cuando la goroutine termine
			// Bucle infinito para procesar trabajos
			for job := range jobs {
				fmt.Printf("Trabajador %d procesando trabajo %d\n", i, job)
			}
		}()
	}

	// Enviamos 100 trabajos al canal de trabajos
	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs) // Cerramos el canal de trabajos para indicar que no se enviarán más trabajos

	wg.Wait() // Esperamos a que todas las goroutines terminen su ejecución
}
