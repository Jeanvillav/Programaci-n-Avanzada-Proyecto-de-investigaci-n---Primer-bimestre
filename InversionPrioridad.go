package main

import (
	"fmt"
	"sync" // Paquete que proporciona primitivas de sincronización como Mutex y WaitGroup
)

func main() {
	// Crear una instancia de Mutex
	mutex := sync.Mutex{}

	// Definir una función para la tarea de alta prioridad
	highPriorityTask := func() {
		// Adquirir el Mutex antes de ejecutar la sección crítica
		mutex.Lock()
		// Diferir la liberación del Mutex al final de la función
		defer mutex.Unlock()
		fmt.Println("Tarea de alta prioridad en ejecución")
	}

	// Definir una función para la tarea de baja prioridad
	lowPriorityTask := func() {
		// Adquirir el Mutex antes de ejecutar la sección crítica
		mutex.Lock()
		// Diferir la liberación del Mutex al final de la función
		defer mutex.Unlock()
		fmt.Println("Tarea de baja prioridad en ejecución")
	}

	// Crear una instancia de WaitGroup
	wg := sync.WaitGroup{}
	// Agregar dos goroutines a WaitGroup
	wg.Add(2)

	// Iniciar una goroutine para ejecutar la tarea de baja prioridad
	go func() {
		// Diferir la llamada a wg.Done() al final de la goroutine
		defer wg.Done()
		lowPriorityTask()
	}()

	// Iniciar una goroutine para ejecutar la tarea de alta prioridad
	go func() {
		// Diferir la llamada a wg.Done() al final de la goroutine
		defer wg.Done()
		highPriorityTask()
	}()

	// Esperar a que ambas goroutines finalicen
	wg.Wait()
}