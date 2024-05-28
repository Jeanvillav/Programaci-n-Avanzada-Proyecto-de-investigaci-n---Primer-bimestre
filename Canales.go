package main

import "fmt"

func main() {
	// Crear un canal para comunicación
	c := make(chan string)
	// Crear un canal para sincronización
	done := make(chan bool)

	// Goroutine receptora
	go func() {
		msg := <-c
		fmt.Println(msg)
		// Enviar una señal de finalización
		done <- true
	}()

	c <- "hello"

	// Esperar a que la goroutine receptora termine
	<-done
}
